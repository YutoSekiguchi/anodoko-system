<template>
  <div>
    <div class="search-form">
      <v-row class="mt-4 mx-auto" justify="center" style="width: 60%">
        <v-autocomplete
          v-model="comicID"
          :items="comicNameIDs"
          outlined
          dense
          color="orange"
          label="漫画名で検索"
          @change="selectComic()"
        ></v-autocomplete>
      </v-row>
    </div>

    <div>
      <v-row>
        <v-col cols="12">
          <SlideShow :items=items />
        </v-col>
      </v-row>
    </div>

    <v-container>
      <div v-if="$store.state.comic.comicFirstVols.length > 0" style="background-color: white;">
        <div v-if="$store.state.comic.comics.length > 0">
          <div class="mt-12">
            <v-col cols="12" class="mb-4 pt-5 " align="center" style="background-color: rgba(0, 0, 0, 0.8);">
              <h2 class="font-weight-bold white--text">人気漫画一覧</h2>
            </v-col>
          </div>
          <v-carousel :v-model="model" height="auto">
            <v-carousel-item v-for="i in colors.length" :key="i" >
              <v-row align="center" justify="center">
                <div>
                  <v-row no-gutters>
                    <v-col v-for="n in 4" :key="n" cols="6" sm="3">
                      <Card
                        :comicData="$store.state.comic.comics[n - 1 + (i - 1) * 4]"
                        :comicFirstData="$store.state.comic.comicFirstVols[n - 1 + (i - 1) * 4]"
                      />
                    </v-col>
                  </v-row>
                </div>
              </v-row>
            </v-carousel-item>
          </v-carousel>
        </div>
      </div>
    </v-container>
    
  </div>
</template>

<script>
import SlideShow from '../components/home/SlideShow.vue'
import Card from '../components/home/Card.vue'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    SlideShow,
    Card
  },
  data() {
    return{
      comicID: null,
      model: 0,
      colors: ["white", "white"],
      items: [
          {
            src: 'https://cs1.anime.dmkt-sp.jp/anime_kv/img/22/67/5/22675_1_9_8b.png?1553502608000',
            title: "鬼滅の刃",
            comment: "12月5日(日)夜11時15分より毎週日曜全国フジテレビ系列にて遊郭編放送中",
            id: 3,
          },
          {
            src: 'https://cs1.anime.dmkt-sp.jp/anime_kv/img/22/54/3/22543_1_9_8b.png?1551346708000',
            title: "五等分の花嫁",
            comment: "2022年初夏映画公開予定",
            id: 8,
          },
          {
            src: 'https://image.video.dmkt-sp.jp/basic/img/title/10008723_sns.jpg',
            title: "進撃の巨人",
            comment: "2022年1月10日NHK総合テレビにて「The Final Season」放送開始",
            id: 1.
          },
          {
            src: 'https://cs1.anime.dmkt-sp.jp/anime_kv/img/24/29/2/24292_1_9_8b.png?1609119057000',
            title: "呪術廻戦",
            comment: "2021年12月24日(金)「劇場版 呪術廻戦0」公開予定",
            id:2,
          },
        ],
      comicsList: [],
      comicFirstVolList: [],
    }
  },
  methods: {
    ...mapActions('comic',[
      'fetchComics',
      'fetchComicFirstVols',
      'resetComic'
    ]),
    ...mapActions('scene', [
      'resetScene'
    ]),
    ...mapActions('character', [
      'resetCharacter'
    ]),
    ...mapActions('favorite', [
      'resetFavorite'
    ]),
    async selectComic() {
      await this.$router.push(`/search/${this.comicID}`)
    }
  },
  computed: {
    ...mapState('comic', ['comics']),
    comicNameIDs() {
      return this.comics.map((comic) => {
        let obj = {}
        obj.value = comic.ID
        obj.text = comic.Name
        return obj
      })
    },
  },
  async created() {
    await this.resetComic()
    await this.resetScene()
    await this.resetCharacter()
    await this.resetFavorite()
    await this.fetchComics()
    await this.fetchComicFirstVols()
  }
}
</script>