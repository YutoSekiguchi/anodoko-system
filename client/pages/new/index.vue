<template>
  <div class="mt-12">
    <form class="mt-12">
      <v-row>
        <v-col align="center">
          <h2 class="font-weight-bold">シーンの登録</h2>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
          <p class="mb-0">漫画名（必須）</p>
          <v-combobox
            v-model="comicName"
            :items="comicNameIDs"
            outlined
            dense
            color="orange"
            label="漫画の名前"
            hide-selected
            @change="selectComic()"
          ></v-combobox>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
          <p class="mb-0">漫画の巻数（必須）</p>
          <v-combobox
            v-model="comicVolume"
            :items="comicVolumeIDs"
            outlined
            dense
            color="orange"
            label="漫画の巻"
            hide-selected
            @change="selectComicVolume()"
          ></v-combobox>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
        <p class="mb-0">ページを入力してください（必須）</p>
        <v-text-field
          v-model="comicPage"
          outlined
          dense
          color="orange"
          label="ページ（半角）"
          hide-selected
        ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
          <p class="mb-0">シーンの種類（必須）</p>
          <v-combobox
            v-model="comicScene"
            :items="comicScenes"
            outlined
            dense
            color="orange"
            label="シーンの種類"
            hide-selected
            @change="selectComicScene()"
          ></v-combobox>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
        <p class="mb-0">シーンの内容を記入してください（必須）</p>
        <v-textarea
          v-model="comicDetailScene"
          outlined
          dense
          color="orange"
          label="どんなシーン？(ex: 〇〇が□□したシーン)"
          hide-selected
          :height="90"
        ></v-textarea>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
        <p class="mb-0">シーンに対するコメント（任意）</p>
        <v-textarea
          v-model="comicComment"
          outlined
          dense
          color="orange"
          label="コメント"
          hide-selected
          :height="60"
        ></v-textarea>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0">
        <p class="mb-0">このシーンを見た時の感情（任意）</p>
        <v-text-field
          v-model="comicEmotion"
          outlined
          dense
          label="感情"
          hide-selected
        ></v-text-field>
        </v-col>
      </v-row>

      <v-col align="center">
        <v-btn color="primary" @click="onSubmit()">次へ</v-btn>
      </v-col>
    </form>

    <v-row class="mt-3 mb-12">
      <v-col align="center">
        <v-btn color="error" @click="toComic()">漫画の追加</v-btn>
      </v-col>

      <v-col align="center">
        <v-btn color="error" @click="toComicVol()">漫画の巻の追加</v-btn>
      </v-col>
    </v-row> 
  </div>
</template>


<script>
import { mapState, mapActions } from 'vuex'

export default {
  components: {
  },
  data() {
    return {
      comicScenes: [
        "戦闘シーン",
        "名シーン",
        "ギャグシーン",
        "死亡シーン",
        "回想シーン",
        "日常シーン",
        "衝撃シーン",
        "解説シーン",
        "感動シーン",
        "胸キュンシーン",
      ],
      comicName: null,
      comicVolume: null,
      comicPage: null,
      comicScene: null,
      comicDetailScene: null,
      comicComment: null,
      comicEmotion: null,
    }
  },
  methods: {
    ...mapActions('comic', [
      'fetchComics',
      'fetchComicVol',
      'fetchComicVols',
    ]),
    ...mapActions('scene', [
      'postScene'
    ]),
    selectComic() {
      console.log(`IDが${this.comicName.value}の漫画が選択されました`);
      if (this.comicVolume != null)  {
        this.comicVolume = null;
      }
      this.fetchComicVols(this.comicName.value);
    },
    selectComicVolume() {
      console.log(`IDが${this.comicVolume.value}の漫画の巻が選択されました`);
    },
    selectComicScene() {
      console.log(`${this.comicScene}を選択`)
    },
    async onSubmit() {
      if (this.comicName == null) {
        alert("漫画名を入力してください")
      } else if (this.comicVolume == null) {
        alert("漫画の巻数を入力してください")
      } else if (this.comicPage == null) {
        alert("ページを入力してください")
      } else if (this.comicScene == null) {
        alert("シーンの種類を入力してください")
      } else if (this.comicDetailScene == null) {
        alert("シーンの内容を記入してください")
      } else {
        var data = {
          uid: this.$store.state.user.userData.ID,
          cvid: this.comicVolume.value,
          page: Number(this.comicPage),
          scene: this.comicScene,
          emotion: this.comicEmotion,
          detailScene: this.comicDetailScene,
          comment: this.comicComment
        }
        await this.postScene(data)
        this.$router.push(`/new/second/${data.uid}/${data.cvid}/${data.page}/${data.scene}/${this.comicName.value}/${this.comicVolume.text}`)
      }
    },
    toComic() {
      this.$router.push('/new/comic')
    },
    toComicVol() {
      this.$router.push('/new/vol')
    }
  },
  computed: {
    ...mapState('comic',
    [
      'comics',
      'comicVol',
      'comicVols',
    ]),
    comicNameIDs() {
      return this.comics.map((comic) => {
        let obj = {};
        obj.text = comic.Name;
        obj.value = comic.ID;
        return obj;
      });
    },
    comicVolumeIDs() {
      return this.comicVols.map((vol) => {
        let obj = {};
        obj.text = vol.Volume;
        obj.value = vol.ID;
        return obj;
      })
    }
  },
  async created() {
    await this.fetchComics()
  }
}
</script>