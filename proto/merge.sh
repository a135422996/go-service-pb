#!/bin/bash

# 遍历当前目录下所有 .proto 文件
for proto_file in *.proto; do
  echo "生成客户端代码: $proto_file"

  # 正确命令格式：将 proto 文件路径直接跟在命令后
  goctl rpc protoc "$proto_file" \
    --go_out=. \
    --go-grpc_out=. \
    --zrpc_out=. \
    --style=go_zero
done

echo "所有客户端代码生成完成！"