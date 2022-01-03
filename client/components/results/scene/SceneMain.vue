<template>
<div>
  <div v-if="sceneData.length >0">
    <v-card
      outlined
      class="mx-auto mt-4 pb-8 mb-12"
      max-width="700px"
      min-width="300px"
      elevation="9"
    >
      <v-row>
        <v-col align="center">
          <h1 class="book_title" >{{ sceneData[0].Name }}</h1>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="2"></v-col>
        <v-col align="center" class="pb-0" cols="8" v-if="sceneData">
          <h2 class="font-weight-bold">{{ sceneData[0].DetailScene }}</h2>
        </v-col>
        <v-col cols="2">
          <div class="heart" style="border: 2px solid rgba(150, 150, 200, 1); border-radius: 30%; width: 50px;" align="center" @click="likeClick(sceneData[0].ScID)" v-if="(this.$store.state.user.userData.ID)">
            <div v-if="this.whetherFavorite.length > 0">
              <span class="" style="color: yellow; font-size: 30px;">★</span>
            </div>
            <div v-if="this.whetherFavorite.length == 0">
              <span class="" style="color: rgba(200, 200, 200, 0.7); font-size: 30px;">★</span>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="2"></v-col>
        <v-col cols="4" class="mt-3" v-if="sceneData != null">
          <v-img :src="this.sceneData[0].ComicVolImage" max-width="200px" contain @click="moveToEbook()" class="image">
          </v-img>
        </v-col>
        <v-col cols="6" align="center" class="pt-0">
          <h4 class="book_vol"><span style="color: red; font-size: 40px;">{{ sceneData[0].Volume }}</span>巻</h4>
          <h4 class="book_page"><span style="color: red; font-size: 40px;" v-if="sceneData">{{ sceneData[0].Page }}</span>ページ</h4>
          <v-btn class="mt-8" color="success" @click="showComment()">コメントを見る</v-btn>
        </v-col>
      </v-row>
    </v-card>
  </div>
</div>
</template>


<script>
import { mapState, mapActions } from 'vuex'

export default {
  props: {
    sceneData: {
      type: Array
    },
  },
  methods: {
    ...mapActions('favorite', [
      'fetchWhetherFavorite',
      'postFavorite',
      'deleteFavorite'
    ]),
    moveToEbook() {
      open(this.sceneData[0].EbookUrl);
    },
    async likeClick(sceneID) {
      var userID = Number(this.$store.state.user.userData.ID)
      if(this.whetherFavorite.length == 0) {
        var data = {
          uid: userID,
          scid: sceneID
        }
        await this.postFavorite(data)
      } else {
        await this.deleteFavorite([
          userID,
          sceneID
        ])
        
      }
      await this.fetchWhetherFavorite([
        userID,
        sceneID
      ])
    },
    showComment() {
      this.$router.push(`/comments/${this.sceneData[0].CVID}/${this.sceneData[0].Page}`);
    },
  },
  computed: {
    ...mapState('scene', [
      'scenesInfo'
    ]),
    ...mapState('favorite', [
      'whetherFavorite',
    ]),
  },
  async created() {
    if (this.$store.state.user.userData.ID){
      await this.fetchWhetherFavorite([
        Number(this.$store.state.user.userData.ID), 
        this.sceneData[0].ScID
      ])
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
.book_vol {
  margin-top: 20%;
  border-bottom: 1px solid rgba(130, 130, 130, 0.7);
}
.image:hover, .heart:hover {
  cursor: pointer;
}
</style>