<template>
  <div>
    <div v-if="this.scenePages">
      
      <div v-if="this.scenePages.length > 0 && this.comic.length > 0">
        <ResultTitle />

          <ResultMain 
            :comicData="comic"
            :scenePagesData="scenePages"
          />

          <Graph 
            :scenePagesData="scenePages"
          />
        
        <div
          v-if="this.resultElseCount == 0"
        >
          <v-row>
            <v-col align="center">
              <v-btn
                color="primary"
                @click="showElseScenes()"
                class="mb-12"
              >
                <p class="mb-0">その他の候補シーンを表示</p>
              </v-btn>
            </v-col>
          </v-row>
        </div>

        <div
          v-if="this.resultElseCount == 1"
        >
          <ResultElse :scenesData="scenesData" />
        </div>

      </div>
    </div>
    
    <div v-else>
      <h2 align="center" class="mt-12">シーンの候補は見つかりませんでした</h2>
    </div>
  </div>
</template>

<script>
import ResultTitle from '~/components/results/scenes/ResultTitle.vue'
import ResultMain from '~/components/results/scenes/ResultMain.vue'
import Graph from '~/components/results/scenes/Graph.vue'
import ResultElse from '~/components/results/scenes/ResultElse.vue'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    ResultTitle,
    ResultMain,
    Graph,
    ResultElse,
  },
  data() {
    return {
      scenesData: [],
      resultElseCount: 0,
    }
  },
  methods: {
    ...mapActions('comic', [
      'fetchComic',
    ]),
    ...mapActions('scene', [
      'fetchScenePagesAsc',
      'fetchScenePagesDesc',
      'fetchScenesUnderVolList'
    ]),
    showElseScenes() {
      this.resultElseCount = 1
      this.elseScenes()
    },
    elseScenes() {
      for(var data of this.scenePages) {
        for(var d in this.scenesUnderVolList) {
          if(data.Page == this.scenesUnderVolList[d].Page && data.Volume == this.scenesUnderVolList[d].Volume) {
            this.scenesData.push(this.scenesUnderVolList[d])
            break
          }
        }
      }
    }
  },
  computed: {
    ...mapState('comic', [
      'comic',
    ]),
    ...mapState('scene', [
      'scenePages',
      'scenesUnderVolList'
    ]),
  },
  async created() {
    await this.fetchComic(this.$route.params.cid)
    
    if(this.$route.params.famous != '名シーン'){
    await this.fetchScenePagesAsc([this.$route.params.chid1, this.$route.params.chid2, this.$route.params.chid3, this.$route.params.chid4, this.$route.params.chid5, this.$route.params.cid, this.$route.params.volume, this.$route.params.scene])
    } else {
    await this.fetchScenePagesDesc([this.$route.params.chid1, this.$route.params.chid2, this.$route.params.chid3, this.$route.params.chid4, this.$route.params.chid5, this.$route.params.cid, this.$route.params.volume, this.$route.params.scene])
    }

    await this.fetchScenesUnderVolList([this.$route.params.chid1, this.$route.params.chid2, this.$route.params.chid3, this.$route.params.chid4, this.$route.params.chid5, this.$route.params.cid, this.$route.params.volume, this.$route.params.scene])
  },
}
</script>