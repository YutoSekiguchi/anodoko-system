<template>
  <div>
    <v-card
        outlined
        class="mx-auto my-12"
        max-width="60%"
        min-width="300px"
        max-height="300px"
        min-height="250px"
        elevation="9"
        >
        <v-col cols="12" class="px-0 py-0 mt-3 mb-4">
          <h2 align="center" class="font-weight-bold" style="border-bottom: 2px solid orange;">{{ this.comicData[0]["Name"] }}</h2>
        </v-col>
        <v-col cols="12" class="px-0 py-0 mt-12 mb-4">
          <h5 align="center" class="font-weight-bold">何巻まで読みましたか？</h5>
        </v-col>
      <v-row>
        <v-col cols="3"></v-col>
        <v-col cols="6" class="px-0 py-0 mt-7">
          <v-select
          v-model="selectedVol"
          :items="comicVolNumbers"
          label="巻数を選択"
          color="orange"
          single-line
          @change="selectVol()"
          ></v-select>
        </v-col>
        <v-col cols="3"></v-col>
      </v-row>
      </v-card>
  </div>
</template>


<script>
export default {
  data() {
    return {
      selectedVol: null
    }
  },
  props: {
    comicVolData: {
      type: Array
    },
    comicData: {
      type: Array
    }
  },
  computed: {
    comicVolNumbers() {
      return this.comicVolData.map((comic) => {
        let obj = {};
        obj.text = comic.Volume;
        obj.value = comic.Volume;
        return obj;
      });
    },
  },
  methods: {
    selectVol() {
      if(this.selectedVol != null){
        this.$router.push(`/search/second/${this.$route.params.cid}/${this.selectedVol}`);
      } else {
        alert('[Error] 再読み込みを行なってください．')
      }
    },
  }
}
</script>