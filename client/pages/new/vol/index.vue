<template>
  <div class="mt-12">
    <form class="mt-12">
      <v-row>
        <v-col align="center">
          <h2 class="font-weight-bold">漫画の巻の登録</h2>
        </v-col>
      </v-row>

      <div class="question">
        <v-col class="px-0 py-0 mt-12 mb-10" cols="12">
          <h5 class="font-weight-bold" align="center">漫画の巻の追加<br>（全ての項目の記入必須）</h5>
        </v-col>

        <v-row class="my-6">
          <v-col cols="3"></v-col>
          <v-col class="px-0 py-0" cols="6">
            <p class="mb-0">漫画名</p>
            <v-combobox
              v-model="comicName"
              :items="comicNameIDs"
              outlined
              dense
              label="漫画の名前"
              hide-selected
              color="orange"
              @change="selectComic()"
            ></v-combobox>
          </v-col>
        </v-row>

        <v-row class="mb-6">
          <v-col cols="3"></v-col>
          <v-col class="px-0 py-0" cols="6">
            <v-text-field
              outlined
              v-model="addComicVolume"
              color="orange"
              dense
              label="何巻？（半角数字）"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row class="mb-6">
          <v-col cols="3"></v-col>
          <v-col class="px-0 py-0" cols="6">
            <v-text-field
              outlined
              v-model="addComicAllPage"
              color="orange"
              dense
              label="総ページ数（半角数字）"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row class="mb-6">
          <v-col cols="3"></v-col>
          <v-col class="px-0 py-0" cols="6">
            <v-text-field
              outlined
              v-model="addEbookUrl"
              color="orange"
              dense
              label="ebookjapanのurl"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row class="mb-6">
          <v-col cols="3"></v-col>
          <v-col class="px-0 py-0" cols="6">
            <v-text-field
              outlined
              v-model="addComicImage"
              color="orange"
              dense
              label="登録する巻の画像のurl"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="3"></v-col>
          <v-col align="center" cols="3">
            <v-btn color="grey" @click="onBack()">戻る</v-btn>
          </v-col>
          <v-col align="center" cols="3">
            <v-btn color="primary" @click="onSubmit()">追加</v-btn>
          </v-col>
        </v-row>
        

      </div>
    </form>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  data() {
    return {
      comicName: null,
      addComicVolume: null,
      addComicAllPage: null,
      addEbookUrl: null,
      addComicImage: null
    }
  },
  methods: {
    ...mapActions('comic', [
      'fetchComics',
      'postComicVol',
    ]),
    selectComic() {
      if (this.comicVolume != null) {
        this.comicVolume = null
      }
    },
    async onSubmit() {
      if (this.comicName == null) {
        alert("漫画名を入力してください")
      } else if (this.addComicVolume == null) {
        alert("巻数を入力してください")
      } else if (this.addComicAllPage == null) {
        alert("総ページ数を入力してください")
      } else if (this.addEbookUrl == null) {
        alert("ebookjapanのurlを入力してください")
      } else if (this.addComicImage == null) {
        alert("漫画の巻の画像のurlを入力してください")
      } else {
        var data = {
          volume: Number(this.addComicVolume),
          cid: this.comicName.value,
          allPage: Number(this.addComicAllPage),
          ebookUrl: this.addEbookUrl,
          comicVolImage: this.addComicImage
        }
        await this.postComicVol(data)
        this.$router.push("/new")
      }
    },
    onBack() {
      this.$router.push("/new")
    }
  },
  computed: {
    ...mapState('comic', [
      'comics'
    ]),
    comicNameIDs() {
      return this.comics.map((comic) => {
        let obj = {}
        obj.text = comic.Name
        obj.value = comic.ID
        return obj
      })
    }
  },
  async created() {
    await this.fetchComics()
  }
}
</script>