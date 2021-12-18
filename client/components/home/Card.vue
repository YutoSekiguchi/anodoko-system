<template>
  <div v-if="comicData">
    <v-hover v-slot="{ hover }">
      <v-col cols="12">
        <v-card
        :loading="loading"
        class="py-6"
        width="350px"
        color="gray"
        outlined
        elevation="11"
        @click="toSearch()"
        >
          <template slot="progress">
            <v-progress-linear
              color="deep-purple"
              indeterminate
            ></v-progress-linear>
          </template>
          
          
          <v-img
            height="200px"
            contain
            :src=comicFirstData.ComicVolImage
          >
          <v-expand-x-transition v-if="hover">
            <div
              class="v-card--reveal text-h5 white--text font-weight-bold"
              style="height: 100%; background-color: rgba(10, 20, 50, 0.8);"
              align="center"      
            >
              <p>{{ comicData.Name }}</p>
              <v-img
                height="110px"
                contain
                :src=comicFirstData.ComicVolImage
              ></v-img>
            </div>
        </v-expand-x-transition>
          </v-img>
        </v-card>
      </v-col>
    </v-hover>
  </div>
  
</template>

<script>
  export default {
    data: () => ({
      hover: false,
      loading: false,
      selection: 1,
    }),
    props:{
        comicData:{
            type: Object
        },
        comicFirstData:{
          type: Object
        },
    },

    methods: {
      toSearch(){
        this.$router.push(`/search/${this.comicData.ID}`);
      },
      reserve () {
        this.loading = true

        setTimeout(() => (this.loading = false), 2000)
      },
    },
  }
</script>
