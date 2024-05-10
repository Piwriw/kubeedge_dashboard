#第一阶段构建
FROM node:lts as builder
WORKDIR /app/ec-dashboard-web
#拷贝源码
COPY . .
#安装依赖
RUN npm install --registry=https://registry.npm.taobao.org

# 开始构建
RUN npm run build:prod

# 第二阶段构建
FROM nginx:alpine
#接口访问地址
ENV EC_DASHBOARD_HOST http://192.168.10.55:8000/
COPY --from=builder /app/ec-dashboard-web/dist/ /usr/share/nginx/html/
COPY default.conf.template /etc/nginx/templates/default.conf.template
EXPOSE 80
