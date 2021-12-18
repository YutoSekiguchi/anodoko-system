<template>
  <div class="main-visual">
    <v-carousel
      hide-delimiters
      cycle
      height=300
      :show-arrows="false"
      class="carousel"
    >
      <v-carousel-item
        v-for="(item, i) in items"
        :key="i"
        @click="nextPage(item)"
        class="carousel-item"
      >
        <v-row>
          <v-col cols="7" align="center" class="px-0">
            <v-img 
              :src="item.src"
              contain
              height="300"
              aspect-ratio="1.7"
            ></v-img>
          </v-col>

          <v-col cols="5" align="center" class="px-0">
            <div class="text-background mr-6">
              <h2 class="font-weight-bold pt-7 mb-8 pb-0 carousel-title">{{ item.title }}</h2>
              <div
                class="mt-0 pt-0 px-2 pb-3"
                v-if="item.comment"
              >
                <p class="carousel-text pb-3">{{ item.comment }}</p>
              </div>
            </div>
          </v-col>
        </v-row>
      </v-carousel-item>
    </v-carousel>
  </div>
</template>


<script>
export default {
  data() {
    return {
    }
  },
  props: {
    items: {
      type: Array
    }
  },
  methods: {
    nextPage(item) {
      if (item.id !== null) {
        this.$router.push(`/search/${item.id}`)
      }
    }
  }
}
</script>


<style scoped>
.carousel {
  background-color: rgba(50,30,0,0.6);
  margin: 20px 0;
}
.text-background {
  margin-top: 70px;
  background-color: rgba(30, 30, 30, 0.9);
}
.main-visual {
  background-image: url("https://www.rakumachi.jp/news/wp-content/uploads/2016/09/iStock-900301626.jpg");
  background-size: cover;
}
.carousel-item:hover {
  cursor: pointer;
}
.carousel-title .carousel-text {
  color: white;
  font-family: sans-serif;
}
</style>