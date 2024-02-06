# etcd-controller
k8s crd controller 测试

两种方式生成自定义代码：
1.默认MODULE
go mod init

/Users/admin/GoProjects/src/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
etcd-controller/pkg/generated \
etcd-controller/pkg/apis \
etcd:v1beta2 \
--output-base ./.. \
--go-header-file hack/boilerplate.go.txt

go mod tidy

将go.mod文件默认MODULE名，rename为github.com/zhangliweixyz/etcd-controller
go mod tidy

2.指定MODULE
go mod init github.com/zhangliweixyz/etcd-controller

/Users/admin/GoProjects/src/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
github.com/zhangliweixyz/etcd-controller/pkg/generated \
github.com/zhangliweixyz/etcd-controller/pkg/apis \
etcd:v1beta2 \
--go-header-file hack/boilerplate.go.txt

go mod tidy

会在当前路径生成github目录，将zz文件和generated目录分别移动到pkg目录下，然后删除github目录即可。
go mod tidy

参数说明：
generate-groups.sh 生成器列表 输出目录 基础包名 组和版本 输出基础路径 版权信息头
