<template>
  <div>
    <ResultElseTitle />
    <v-row>
      <v-col v-for="(items, i) in scenesData.slice(1)" :key="i" cols="12" sm="6" md="4">
        <div>
          <v-col align="center">
            <SceneCard :comicData="items" />
          </v-col>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import ResultElseTitle from '~/components/results/character/ResultElseTitle.vue'
import SceneCard from '~/components/results/character/SceneCard.vue'

export default {
  components: {
    ResultElseTitle,
    SceneCard
  },
  props: {
    scenesData: {
      type: Array
    }
  },
}
</script>