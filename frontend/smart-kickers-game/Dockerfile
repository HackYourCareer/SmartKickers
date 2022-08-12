FROM node:alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm i
#RUN npm install react-scripts -g --silent
COPY . ./
RUN npm run build

FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
