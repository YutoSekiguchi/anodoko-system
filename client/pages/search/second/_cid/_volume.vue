<template>
  <div class="searchsecond">
    <v-col>

      <SearchTitle />

      <div
        v-if="comic.length > 0">
        <SearchFormDetail 
          :comicData="comic"
          :comicUnderVolsData="comicUnderVols"
          :characterUnderVolsData="characterUnderVols"
        />
      </div>
      
    </v-col>
  </div>
</template>


<script>
import SearchTitle from '~/components/search/SearchTitle.vue'
import SearchFormDetail from '~/components/search/second/SearchFormDetail.vue'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    SearchTitle,
    SearchFormDetail,
  },
  methods: {
    ...mapActions('comic', [
      'fetchComic',
      'fetchComicUnderVols',
    ]),
    ...mapActions('character', [
      'fetchCharacterUnderVols'
    ]),
    
  },
  computed: {
    ...mapState('comic', [
      'comic',
      'comicUnderVols'
    ]),
    ...mapState('character', [
      'characterUnderVols',
    ]),
  },
  async created() {
    await this.fetchComic(this.$route.params.cid)
    await this.fetchComicUnderVols([this.$route.params.cid, this.$route.params.volume])
    await this.fetchCharacterUnderVols([this.$route.params.cid, this.$route.params.volume])
  },
}
</script>