package controllers

import "github.com/kataras/iris/context"

var clusterController ClusterController

type ClusterController struct {
	BaseAction
}

func ClusterListHandler(ctx context.Context) {
	clusterController.GetList(ctx)
}

func ClusterSaveHandler(ctx context.Context) {
	clusterController.Save(ctx)
}
