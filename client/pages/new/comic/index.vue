<template>
  <div class="mt-12">
    <form class="mt-12">
      <v-row>
        <v-col align="center">
          <h2 class="font-weight-bold">漫画の登録</h2>
        </v-col>
      </v-row>

      <div class="question">
        <v-col class="px-0 py-0 mt-12 mb-10" cols="12">
          <h5 class="font-weight-bold" align="center">漫画の追加</h5>
        </v-col>

        <v-row class="mb-6">
          <v-col cols="2"></v-col>
          <v-col class="px-0 py-0" cols="8">
            <v-text-field
              outlined
              v-model="addComic"
              label="追加したい漫画名の入力"
            ></v-text-field>
          </v-col>
          <v-col cols="2"></v-col>
        </v-row>
        <v-col align="center">
          <v-btn color="primary" @click="onSubmit()">追加</v-btn>
        </v-col>
      </div>

    </form>
  </div>
</template>


<script>
import { mapActions } from 'vuex'

export default {
  data() {
    return {
      addComic: null
    }
  },
  methods: {
    ...mapActions('comic', [
      'postComic',
    ]),
    async onSubmit() {
      var data = {
        name: this.addComic
      }
      await this.postComic(data)
      this.$router.push("/new")
    }
  }
}
</script>