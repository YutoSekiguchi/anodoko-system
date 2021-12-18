<template>
  <div v-if="comicData">

    <v-row>
      <v-card
        outlined
        class="mx-auto mt-4 pb-8 mb-4 "
        width="300px"
        elevation="9"
      >
        <v-row>
          <v-col align="center">
            <h4 class="book_title" >{{ this.$store.state.comic.comic[0].Name }}</h4>
          </v-col>
        </v-row>
        <v-row>
          <v-col align="center" class="pb-0">
            <h5 class="font-weight-bold">{{ comicData.DetailScene }}</h5>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="2"></v-col>
          <v-col cols="4" class="mt-3"><v-img :src="comicData.ComicVolImage" max-width="200px" contain @click="moveToEbook()" class="image">
          </v-img></v-col>
          <v-col cols="6" align="center">
            <h6 class="book_vol"><span style="color: red; font-size: 20px;">{{ comicData.Volume }}</span>巻</h6>
            <h6 class="book_page"><span style="color: red; font-size: 20px;">{{ comicData.Page }}</span>ページ</h6> 
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-btn class="" color="success" @click="showComment()">コメントを閲覧する</v-btn>
          </v-col>
        </v-row>
      </v-card>
    </v-row>
    
  </div>
  
</template>

<script>
  export default {
    props:{
      comicData:{
        type: Object
      }
    },

    methods: {
      moveToEbook() {
        open(this.comicData.EbookUrl);
      },
      showComment() {
        this.$router.push(`/comments/${this.comicData.CVID}/${this.comicData.Page}`);
      }
    },
  }
</script>

<style scoped>
.book_title {
  font-weight: bold;
  color: rgb(255,170,0);
  color: white;
  padding: 20px 0;
  border-bottom: 2px solid gray;
  background-color: salmon;
}

.top-text {
  color: rgb(165, 91, 42);
}

.question-mark {
  font-style: italic;
  font-size: 60px;
  color: orange;
}

.book_vol {
  margin-top: 20%;
  border-bottom: 1px solid rgba(130, 130, 130, 0.7);
}
.image:hover, .heart:hover {
  cursor: pointer;
}
.small {
    margin:  150px auto;
  }
</style>