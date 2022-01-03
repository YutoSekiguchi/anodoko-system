<template>
  <div class="mt-12">
    <form class="mt-12">
      <v-row>
        <v-col align="center">
          <h2 class="font-weight-bold">シーンの登録</h2>
        </v-col>
      </v-row>

      <div class="qustion">
        <v-col class="px-0 py-0 mt-12 mb-10" cols="12" >
          <h3 class="font-weight-bold" align="center">登場キャラクターを選択<br>(残り<span style="color: red; font-size: 25px;">{{ 5 - this.selectedCharacters.length }}</span>キャラ登録可能)
          </h3>
        </v-col>

        <v-row class="">
          <v-col cols="2"></v-col>
          <v-col cols="8" class="px-0 py-0">
            <v-autocomplete
              auto-select-first
              clearable
              multiple
              solo
              color="orange"
              v-model="selectedCharacters"
              :items="comicVolCharacters"
              label="キャラクターを選択"
              @change="showCharacters()"
            ></v-autocomplete>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>
        <v-col align="center" class="pt-0 mb-12">
          <v-btn color="primary" @click="onSubmit()">送信</v-btn>
        </v-col>
      </div>


      <div style="border: 1px solid yellow; border-radius: 50px;" class="mx-12">
        <v-col cols="12" class="px-0 py-0 mt-12 mb-10">
          <h5 align="center" class="font-weight-bold"><span class="" style="font-size: 20px; color: lightblue;">上にキャラクターがなかった場合…</span><br><br>登場キャラクターを新規追加</h5>
        </v-col>
        <v-row class="mb-6">
          <v-col cols="2"></v-col>
          <v-col cols="8" class="px-0 py-0">
            <v-text-field
              outlined
              v-model="newCharacter1"
              color="orange"
              label="追加したいキャラクター名を入力"
            ></v-text-field>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>

        <v-col align="center">
          <v-btn color="error" @click="onAdd()">新規キャラの追加</v-btn>
        </v-col>

      </div>   
    </form>
  </div>
</template>


<script>
import { mapState, mapActions } from 'vuex'

export default {
  data() {
    return {
      selectedCharacters: [],
      newCharacter1: null,
    }
  },
  methods: {
    ...mapActions('comic', [
      'fetchComics',
    ]),
    ...mapActions('character', [
      'fetchCharacterUnderVols',
      'fetchCharacterId',
      'postAppearanceCharacter',
      'postCharacter'
    ]),
    ...mapActions('scene', [
      'fetchUserScenesInfo'
    ]),
    async onSubmit() {
      for(var i of this.selectedCharacters) {
        let data = {
          scid: this.userScenesInfo[0].ID,
          cvid: this.userScenesInfo[0].CVID,
          chid: i
        }
        await this.postAppearanceCharacter(data)
      }
      this.$router.push("/new")
    },
    async onAdd() {
      if (this.newCharacter1 != null) {
        let data = {
          name: this.newCharacter1
        }
        await this.postCharacter(data)
        alert(`${this.newCharacter1}を追加しました`)
        await this.fetchCharacterId(data.name)
        await this.selectedCharacters.push(this.characterId[0].ID)
        this.newCharacter1 = null
      }
    },
    showCharacters() {
      console.log(this.selectedCharacters);
    },
  },
  computed: {
    ...mapState('comic', [
      'comics',
    ]),
    ...mapState('character', [
      'characterUnderVols',
      'characterId',
    ]),
    ...mapState('scene', [
      'userScenesInfo',
    ]),
    comicVolCharacters() {
      if (this.characterUnderVols != null) {
        return this.characterUnderVols.map((character) => {
          let obj = {}
          obj.text = character.Name
          obj.value = character.Chid
          return obj
        })
      }
    }
  },
  async created() {
    await this.fetchComics()
    await this.fetchCharacterUnderVols([this.$route.params.cid, this.$route.params.volume])
    await this.fetchUserScenesInfo([this.$route.params.uid, this.$route.params.cvid, this.$route.params.page, this.$route.params.scene])
  }
}
</script>