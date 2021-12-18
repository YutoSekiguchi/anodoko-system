<template>
  <div>
    <v-card
      outlined
      class="mx-auto my-12"
      max-width="60%"
      min-width="300px"
      elevation="9"
    >
      <v-col cols="12" class="px-0 py-0 mt-3 mb-4">
        <h1 align="center" class="font-weight-bold pb-3" style="border-bottom: 4px solid orange;">{{ comicData[0].Name }}</h1>
      </v-col>

      <div class="question">
        <v-col cols="12" class="px-0 py-0 mt-12 mb-10">
          <h4 align="center" class="font-weight-bold">登場キャラクターを選択<br>（5キャラ以内）</h4>
        </v-col>
        <v-row class="mb-6">
          <v-col cols="2"></v-col>
          <v-col cols="8" class="px-0 py-0">
            <v-autocomplete
              auto-select-first
              clearable
              multiple
              solo
              outlined
              color="orange"
              v-model="selectedCharacters"
              :items="comicVolCharacters"
              label="キャラクターを選択"
            ></v-autocomplete>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>
      </div>

      <div class="question">
        <v-col cols="12" class="px-0 py-0 mt-8 mb-8">
          <h4 align="center" class="font-weight-bold">どんなシーン？</h4>
        </v-col>
        <v-row class="mb-6">
          <v-col cols="2"></v-col>
          <v-col cols="8" class="px-0 py-0">
            <v-select
              solo
              clearable
              outlined
              color="orange"
              v-model="selectedScene"
              :items="kindOfScenes"
              label="シーンを選択"
            ></v-select>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>
      </div>

      <div class="question">
        <v-col cols="12" class="px-0 py-0 mt-8 mb-8">
          <h4 align="center" class="font-weight-bold">名シーン？マニアック？</h4>
        </v-col>
        <v-row class="mb-6">
          <v-col cols="2"></v-col>
          <v-col cols="8" class="px-0 py-0">
            <v-select
              solo
              clearable
              outlined
              color="orange"
              v-model="selectedFamous"
              :items="famousOrManiac"
              label="選択してください"
            ></v-select>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>
      </div>
      

      <div class="button my-6" align="center">
        <v-btn class="" color="primary" @click="searchSubmit()">検索</v-btn>
      </div>
    </v-card>

      <v-row>
        <v-col cols="2"></v-col>
        <v-col cols="4" align="center">
          <BackHomeButton />
        </v-col>
        <v-col cols="4" align="center">
          <BackSearchButton />
        </v-col>
        <v-col cols="2"></v-col>
    </v-row>
  </div>
  
</template>

<script>
import BackHomeButton from "~/components/common/BackHomeButton.vue"
import BackSearchButton from "~/components/search/second/BackSearchButton.vue"

export default {
  components: {
    BackHomeButton,
    BackSearchButton
  },
  data() {
    return {
      selectedCharacters: [],
      selectedScene: null,
      selectedFamous: null,
      famousScene: true,
      kindOfScenes: [
        "戦闘シーン",
        "衝撃シーン",
        "名シーン",
        "感動シーン",
        "ギャグシーン",
        "死亡シーン",
        "回想シーン",
        "日常シーン",
        "胸キュンシーン",
      ],
      famousOrManiac: [
        "名シーン",
        "マニアックシーン"
      ]
    }
  },
  props: {
    comicData: {
      type: Array
    },
    comicUnderVolsData: {
      type: Array
    },
    characterUnderVolsData: {
      type: Array
    }
  },
  methods: {
    searchSubmit() {
        if (this.selectedCharacters.length == 0){
          alert("登場キャラクターを選択してください");
        } else if (this.selectedCharacters.length > 5){
          alert("登場キャラクターは5人以下にしてください");
        } else if (this.selectedScene == null) {
          alert("どんなシーンだったか入力してください");
        } else if (this.selectedFamous == null) {
          alert("名シーンかマニアックか答えてください");
        }
        else {
          var len = this.selectedCharacters.length;
          if (len < 5) {
            for(var i = 0; i < 5 - len; i++) {
              this.selectedCharacters.push(0);
            }
          }
          if (len > 5){
            alert("エラーが発生しました．再読み込みしてください");
          }
          if(this.$store.getters['user/isLoggedIn']){
            this.$router.push(`/results/characters/${this.selectedCharacters[0]}/${this.selectedCharacters[1]}/${this.selectedCharacters[2]}/${this.selectedCharacters[3]}/${this.selectedCharacters[4]}/${this.$route.params.cid}/${this.$route.params.volume}/${this.selectedScene}/${this.selectedFamous}/${this.$store.state.user.userData.ID}`);
          }else{
            this.$router.push(`/results/characters/${this.selectedCharacters[0]}/${this.selectedCharacters[1]}/${this.selectedCharacters[2]}/${this.selectedCharacters[3]}/${this.selectedCharacters[4]}/${this.$route.params.cid}/${this.$route.params.volume}/${this.selectedScene}/${this.selectedFamous}`);
          }
        }
    }
  },
  computed: {
    comicVolCharacters() {
      return this.characterUnderVolsData.map((character) => {
        let obj = {}
        obj.text = character.Name
        obj.value = character.Chid
        return obj
      })
    }
  }
}
</script>