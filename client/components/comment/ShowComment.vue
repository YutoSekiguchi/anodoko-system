<template>
  <div v-if="commentData.Comment != ''">
    <v-row>
      <v-col cols="1" sm="3"></v-col>
      <v-col cols="10" sm="6">
        <v-card class="mt-10" color=""  elevation="10">
          <v-row class="ml-2">
            <v-img :src="commentData.Image" max-width="50px" style="border-radius: 50%;"></v-img>
            <h6 class="font-weight-bold ml-3" style="line-height: 50px;">{{ commentData.Name }}<span class="ml-3" style="color: gray; font-size: 12px;">{{ commentData.CreatedAt.substr(0, commentData.CreatedAt.indexOf('T')) }}</span></h6>
          </v-row>
          <v-row>
            <v-col class="pt-0 mb-2">
              <h4 class="font-weight-bold ml-12">{{ commentData.Comment }}</h4>
            </v-col>
          </v-row>
          <hr>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  props: {
    commentData: {
      type: Object
    }
  }
}
</script>