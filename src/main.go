package main

import (
	"os/exec"
	"log"
	"flag"
	"strings"
)

type ObjectSpec struct{
	Kind, Namespace string
}

func FindObjects(obj ObjectSpec) {

	var (
		instruct string
		out []string
	)
	
	ObjectNameList := []string{}


	if obj.Namespace == "" {
		instruct = "kubectl get " +  obj.Kind + " -o=name"
	} else {
		instruct = "kubectl get " +  obj.Kind + " -n " + obj.Namespace + " -o=name"
	}

	cmd := exec.Command("sh", "-c", instruct)
	res, _ := cmd.Output()
	out = strings.Fields(string(res))

	for ind:=0; ind<len(out); ind+=1 {
		name := strings.Split(out[ind], "/")[1]
		ObjectNameList = append(ObjectNameList, name)
	}

	DeleteObjects(obj, ObjectNameList)
}

func DeleteObjects(obj ObjectSpec, ObjectNameList []string) {

	var instruct string

	for _, name := range ObjectNameList {
		if obj.Namespace == "" {
			instruct = "kubectl delete " +  obj.Kind + " " + name
		} else {
			instruct = "kubectl delete " +  obj.Kind + " " + name + " -n " + obj.Namespace 
		}

		cmd := exec.Command("sh", "-c", instruct)
		res, _ := cmd.Output()
		log.Printf("%s", res)
	}
}


func main(){
	
	var (
		Kind, Namespace string
	)

	flag.StringVar(&Kind, "kind", "pods", "Kubectl Object Kind to be removed")
	flag.StringVar(&Namespace, "namespace", "", "Namespace where the object to be removed resides")
	flag.Parse()

	obj := ObjectSpec{
		Kind: Kind, 
		Namespace: Namespace,
	}

	FindObjects(obj)
}
