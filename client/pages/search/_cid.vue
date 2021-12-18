<template>
  <div>
    <SearchTitle />
    <v-container>
      
      <div
        v-if="$store.state.comic.comicVols.length > 0"
      >
        <SearchFormVol
          :comicData="$store.state.comic.comic"
          :comicVolData="$store.state.comic.comicVols" />
      </div>
      <BackHomeButton />
    </v-container>
  </div>
</template>


<script>
import BackHomeButton from '~/components/common/BackHomeButton.vue'
import SearchFormVol from '~/components/search/SearchFormVol.vue'
import SearchTitle from '~/components/search/SearchTitle.vue'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    BackHomeButton,
    SearchFormVol,
    SearchTitle
  },
  methods: {
    ...mapActions('comic', [
      'fetchComic',
      'fetchComicVols',
    ]),
  },
  computed: {
    ...mapState('comic', [
      'comic',
      'comicVols',
    ]),
  },
  async created() {
    await this.fetchComic(this.$route.params.cid)
    await this.fetchComicVols(this.$route.params.cid)
  },
}
</script>