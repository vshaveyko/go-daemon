package process

import (
	"github.com/jinzhu/gorm"
	"github.com/vshaveyko/test-go-daemon/utils"
)

type Pipeline struct {
	gorm.Model
	Connectors []Connector
}

func (pipe *Pipeline) toShellCmd() (res []string) {

	return []string{
		"luigi",
		"--module main WrapperTask",
		"--depends-on",
		"'" + pipe.toJson() + "'",
	}

}

type interfacemap map[string]interface{}
type dependencyMap map[int][]EndPoint
type targetMap map[int]EndPoint

func (pipe *Pipeline) toJson() string {
	res := pipe.getFlow()

	return utils.Jsonify(res)
}

func (pipe *Pipeline) getFlow() []interface{} {

	dependencyMap := dependencyMap{}
	targetMap := targetMap{}

	for _, c := range pipe.Connectors {
		dependencyMap[c.Target.Id] = append(dependencyMap[c.Target.Id], c.Source)
		targetMap[c.Target.Id] = c.Target
	}

	var res []interface{}

	for _, v := range resolveDependencyMap(dependencyMap, targetMap) {
		res = append(res, v)
	}

	return res

}

func resolveDependencyMap(dependencyMap dependencyMap, targetMap targetMap) map[int]interfacemap {

	var resolveDependency func(EndPoint) interfacemap
	endPointData := make(map[int]interfacemap)

	resolveDependency = func(target EndPoint) interfacemap {
		if v, ok := endPointData[target.Id]; ok {
			return v
		}

		datum := make(interfacemap)

		for k, v := range target.toMap() {
			datum[k] = v
		}

		depends_on := make([]interfacemap, 0)

		for _, ep := range dependencyMap[target.Id] {
			depends_on = append(depends_on, resolveDependency(ep))
			delete(endPointData, ep.Id)
		}

		datum["depends_on"] = depends_on

		endPointData[target.Id] = datum

		return endPointData[target.Id]
	}

	for epId, _ := range dependencyMap {
		resolveDependency(targetMap[epId])
	}

	return endPointData

}
