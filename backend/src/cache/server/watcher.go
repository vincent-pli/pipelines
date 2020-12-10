package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/kubeflow/pipelines/backend/src/cache/client"
	"github.com/kubeflow/pipelines/backend/src/cache/model"
	"github.com/peterhellberg/duration"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/termination"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"knative.dev/pkg/logging"
)

const (
	ArgoCompleteLabelKey   string = "workflows.argoproj.io/completed"
	MetadataExecutionIDKey string = "pipelines.kubeflow.org/metadata_execution_id"
	MaxCacheStalenessKey   string = "pipelines.kubeflow.org/max_cache_staleness"
)

func WatchPods(namespaceToWatch string, clientManager ClientManagerInterface) {
	k8sCore := clientManager.KubernetesCoreClient()

	for {
		listOptions := metav1.ListOptions{
			Watch:         true,
			LabelSelector: CacheIDLabelKey,
		}
		watcher, err := k8sCore.PodClient(namespaceToWatch).Watch(listOptions)

		if err != nil {
			log.Printf("Watcher error:" + err.Error())
		}

		for event := range watcher.ResultChan() {
			pod := reflect.ValueOf(event.Object).Interface().(*corev1.Pod)
			if event.Type == watch.Error {
				continue
			}
			log.Printf((*pod).GetName())

			if !isPodCompletedAndSucceeded(pod) {
				log.Printf("Pod %s is not completed or not in successful status.", pod.ObjectMeta.Name)
				continue
			}
			log.Println("---- step - 1")
			if isCacheWriten(pod.ObjectMeta.Labels) {
				continue
			}
			log.Println("---- step - 2")
			executionKey, exists := pod.ObjectMeta.Annotations[ExecutionKey]
			if !exists {
				continue
			}
			log.Println("---- step - 3")
			// executionOutput, exists := pod.ObjectMeta.Annotations[ArgoWorkflowOutputs]
			executionOutput, err := parseResult(pod)
			if err != nil {
				log.Printf("Result of Pod %s not parse success.", pod.ObjectMeta.Name)
				continue
			}
			log.Println("---- step - 4")
			executionOutputMap := make(map[string]interface{})
			executionOutputMap[TektonTaskrunOutputs] = executionOutput
			executionOutputMap[MetadataExecutionIDKey] = pod.ObjectMeta.Labels[MetadataExecutionIDKey]
			executionOutputJSON, _ := json.Marshal(executionOutputMap)

			executionMaxCacheStaleness, exists := pod.ObjectMeta.Annotations[MaxCacheStalenessKey]
			var maxCacheStalenessInSeconds int64 = -1
			if exists {
				maxCacheStalenessInSeconds = getMaxCacheStaleness(executionMaxCacheStaleness)
			}

			executionTemplate := pod.ObjectMeta.Annotations[TektonTaskrunTemplate]
			executionToPersist := model.ExecutionCache{
				ExecutionCacheKey: executionKey,
				ExecutionTemplate: executionTemplate,
				ExecutionOutput:   string(executionOutputJSON),
				MaxCacheStaleness: maxCacheStalenessInSeconds,
			}

			cacheEntryCreated, err := clientManager.CacheStore().CreateExecutionCache(&executionToPersist)
			if err != nil {
				log.Println("Unable to create cache entry.")
				continue
			}
			log.Println("---- step - 5")
			err = patchCacheID(k8sCore, pod, namespaceToWatch, cacheEntryCreated.ID)
			if err != nil {
				log.Printf(err.Error())
			}
		}
	}
}

func parseResult(pod *corev1.Pod) (string, error) {
	log.Println("Start parse result from pod.")

	logger := logging.FromContext(context.TODO())
	output := []*v1beta1.TaskRunResult{}

	containersState := pod.Status.ContainerStatuses
	if containersState == nil || len(containersState) == 0 {
		return "", fmt.Errorf("No container status found")
	}
	fmt.Printf("+++++++++++++++ %+v", containersState)
	for _, state := range containersState {
		fmt.Println("000000000000000")
		if state.State.Terminated != nil && len(state.State.Terminated.Message) != 0 {
			fmt.Println("111111111")
			msg := state.State.Terminated.Message
			results, err := termination.ParseMessage(logger, msg)
			fmt.Println(msg)
			fmt.Printf("xx %+v", results)
			fmt.Println("222222222")
			if err != nil {
				logger.Errorf("termination message could not be parsed as JSON: %v", err)
				return "", fmt.Errorf("termination message could not be parsed as JSON: %v", err)
			}
			fmt.Println("3333333")
			for _, r := range results {
				if r.ResultType == v1beta1.TaskRunResultType {
					itemRes := v1beta1.TaskRunResult{}
					itemRes.Name = r.Key
					itemRes.Value = r.Value
					output = append(output, &itemRes)
				}
			}
			fmt.Println("4444444")
			// assumption only on step in a task
			break
		}
	}
	fmt.Printf("---- %+v", output)
	b, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	fmt.Println("________________")
	fmt.Println(string(b))
	return string(b), nil
}

func isPodCompletedAndSucceeded(pod *corev1.Pod) bool {
	return pod.Status.Phase == corev1.PodSucceeded
}

func isCacheWriten(labels map[string]string) bool {
	cacheID := labels[CacheIDLabelKey]
	return cacheID != ""
}

func patchCacheID(k8sCore client.KubernetesCoreInterface, podToPatch *corev1.Pod, namespaceToWatch string, id int64) error {
	labels := podToPatch.ObjectMeta.Labels
	labels[CacheIDLabelKey] = strconv.FormatInt(id, 10)
	log.Println(id)
	var patchOps []patchOperation
	patchOps = append(patchOps, patchOperation{
		Op:    OperationTypeAdd,
		Path:  LabelPath,
		Value: labels,
	})
	patchBytes, err := json.Marshal(patchOps)
	if err != nil {
		return fmt.Errorf("Unable to patch cache_id to pod: %s", podToPatch.ObjectMeta.Name)
	}
	_, err = k8sCore.PodClient(namespaceToWatch).Patch(podToPatch.ObjectMeta.Name, types.JSONPatchType, patchBytes)
	if err != nil {
		return err
	}
	log.Printf("Cache id patched.")
	return nil
}

// Convert RFC3339 Duration(Eg. "P1DT30H4S") to int64 seconds.
func getMaxCacheStaleness(maxCacheStaleness string) int64 {
	var seconds int64 = -1
	if d, err := duration.Parse(maxCacheStaleness); err == nil {
		seconds = int64(d / time.Second)
	}
	return seconds
}
