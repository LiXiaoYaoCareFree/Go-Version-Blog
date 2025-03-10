Docker安装ollama
拉取ollama镜像：

docker pull ollama/ollama
如果下载速度太慢，可使用swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ollama/ollama镜像源，即：

docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ollama/ollama:0.3.13
docker tag swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ollama/ollama:0.3.13 ollama/ollama:0.3.13
执行如下命令启动Ollama镜像。使用CPU推理运行：

docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama:0.3.13
该命令将Ollama服务暴露于主机的11434端口。

接下来我们可以尝试进入容器，通过容器内的ollama命令运行llama3.2:3b模型。

命令如下：
docker start ollama
docker exec -it ollama bash
ollama run llama3.2:3b
Ollama运行大模型服务
Ollama通过ollama命令操作大模型。ollama命令支持的参数如下：

root@26fd85bc69b6:/# ollama
Usage:
ollama [flags]
ollama [command]

Available Commands:
serve       Start ollama
create      Create a model from a Modelfile
show        Show information for a model
run         Run a model
stop        Stop a running model
pull        Pull a model from a registry
push        Push a model to a registry
list        List models
ps          List running models
cp          Copy a model
rm          Remove a model
help        Help about any command

Flags:
-h, --help      help for ollama
-v, --version   Show version information

Use "ollama [command] --help" for more information about a command.

所有Ollama支持的大语言模型，可在Ollama仓库中站点检索。链接为：https://ollama.com/library

其中常用的命令如下：

启动并进入一个大模型的交互式问答模式：

ollama run llama3.2:3b
该命令会自动下载llama3.2:3b模型，然后进入到交互式问答模式。如果模型已下载则直接键入交互式问答模式。例如：

root@26fd85bc69b6:/# ollama run llama3.2:3b
>>> who are you
I'm an artificial intelligence model known as Llama. Llama stands for "Large Language Model Meta AI."
键入/bye退出交互式问答模式

列出所有已下载模型：

root@26fd85bc69b6:/# ollama list
NAME                         ID              SIZE      MODIFIED
qwen2.5:7b                   845dbda0ea48    4.7 GB    9 days ago
llama3.2:3b                  a80c4f17acd5    2.0 GB    10 days ago
列出正在运行的模型：

root@26fd85bc69b6:/# ollama ps
NAME           ID              SIZE      PROCESSOR    UNTIL
llama3.2:3b    a80c4f17acd5    3.5 GB    100% CPU     2 minutes from now
让指定的模型停止运行：

ollama stop llama3.2:3b
删除模型：

ollama rm llama3.2:3b


ssh -L 11434:127.0.0.1:11434 root@192.168.179.135
