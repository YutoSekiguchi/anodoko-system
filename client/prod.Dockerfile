FROM node:15.8


WORKDIR /app
ENV NUXT_HOST 0.0.0.0

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

EXPOSE 3000

CMD [ "npm", "run", "start" ]