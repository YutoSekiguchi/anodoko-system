<template>
  <div>
    <div v-if="this.scene">
      <div v-if="this.scene.length > 0">
        <ResultTitle />
        <SceneMain :sceneData="scene" />
      </div>
    </div>
    
    <div v-else>
      <h2 align="center" class="mt-12">シーンは見つかりませんでした。</h2>
    </div>
  </div>
</template>

<script>
import ResultTitle from '~/components/results/scenes/ResultTitle.vue'
import SceneMain from '~/components/results/scene/SceneMain.vue'
import Graph from '~/components/results/scenes/Graph.vue'
import ResultElse from '~/components/results/scenes/ResultElse.vue'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    ResultTitle,
    SceneMain,
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
    ...mapActions('scene', [
      'fetchScene',
    ]),
  },
  computed: {
    ...mapState('scene', [
      'scene',
    ]),
  },
  async created() {
    await this.fetchScene(this.$route.params.scid)
  },
}
</script>