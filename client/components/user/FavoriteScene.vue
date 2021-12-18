<template>
<div>
  <div v-for="(data, i) in this.favoritesList" :key="i">
    <v-row style="border-bottom: 0.5px solid grey;" class="mr-12 mt-3">
      <Scenes :data="data" />
      <ScenesDetail :data="data" />
      <v-col cols="1">
        <v-btn color="error" height="30" @click="deleteLike(data.Id)"><h6>削除</h6></v-btn>
      </v-col>
    </v-row>
  </div>
</div>
</template>

<script>
import Scenes from '~/components/user/Scenes.vue'
import ScenesDetail from '~/components/user/ScenesDetail.vue'
import { mapState, mapActions } from 'vuex'

export default {
  components:{
    Scenes,
    ScenesDetail
  },
  methods: {
    ...mapActions('favorite', [
      'fetchFavoritesList',
      'deleteFavorite'
    ]),
    async deleteLike(scid) {
      await this.deleteFavorite([this.$store.state.user.userData.ID, scid])
      await this.fetchFavoritesList(this.$store.state.user.userData.ID)
    }
  },
  computed: {
    ...mapState('favorite', [
      'favoritesList'
    ])
  },
  async created() {
    await this.fetchFavoritesList(this.$store.state.user.userData.ID)
  }
}
</script>