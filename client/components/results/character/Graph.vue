<template>
  <div class="mb-12 graph-box">
    <v-col align="center">
      <v-btn v-if="clickCount == 0" @click="fillData(comicVol)" class="error">{{ comicVol.Volume }}巻のグラフを表示</v-btn>
      <v-btn v-if="clickCount == 1" @click="hideGraph()" class="error mb-5">グラフの非表示</v-btn>
    </v-col>
    <div class="small py-0 my-0" v-if="clickCount != 0" >
      <LineChart :chart-data="datacollection" class="graph" :height="120"></LineChart>
    </div>
  </div>
</template>

<script>
import LineChart from '~/components/results/LineChart.js'

import { mapState, mapActions } from 'vuex'

export default {
  components: {
    LineChart
  },
  data() {
    return {
      clickCount: 0,
      datacollection: {},
    }
  },
  props: {
    scenePagesData: {
      type: Array
    },
  },
  methods: {
    ...mapActions('comic', [
      'fetchComicVol',
    ]),
    ...mapActions('scene', [
      'fetchScenesVolList',
    ]),
    fillData(d) {
      this.clickCount = 1
      let all_pages = this.comicVol[0].AllPage
      let arr_1 = [...Array(all_pages)].map(() => 0);
      let arr_2 = [...Array(all_pages)].map(() => 0);

      for(var i in this.scenePagesData) {
        if (d[0].Volume == this.scenePagesData[i].Volume && d[0].CID == this.scenePagesData[i].Cid) {
          arr_1[this.scenePagesData[i].Page] = this.scenePagesData[i].Countpage;
        }
      }

      for(var j in this.scenesVolList) {
        if (d[0].Volume == this.scenesVolList[j].Volume && d[0].CID == this.scenesVolList[j].Cid) {
          arr_2[this.scenesVolList[j].Page] = this.scenesVolList[j].Countpage;
        }
      }

      this.datacollection = {
        labels: [...Array(all_pages).keys()].map((d) => {return d}),
        datasets: [
          {
            label: `この巻の探しているシーンに近い人気ページ`,
            borderColor: "rgba(200,30,30, 1)",
            backgroundColor: "rgba(200,70,30, 0.4)",
            data: arr_1,
          },
          {
            label: `この巻の人気ページ`,
            borderColor : 'rgba(22,20,200, 0.8)',//塗りつぶす色
            backgroundColor: "rgba(22,20,200, 0.4)",
            data: arr_2,
          },
        ],
      }
    },
    hideGraph() {
      this.clickCount = 0
    }
  },
  computed: {
    ...mapState('comic', [
      'comicVol',
    ]),
    ...mapState('scene', [
      'scenesVolList',
    ])
  },
  async created() {
    await this.fetchScenesVolList([this.$route.params.cid, this.$store.state.scene.scenePages[0].Volume])
    await this.fetchComicVol([this.$route.params.cid, this.$store.state.scene.scenePages[0].Volume])
  }
}
</script>