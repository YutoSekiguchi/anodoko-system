<template>
  <div>
    <h2 align="center" class="font-weight-bold mt-12" style="color: orangered; border-bottom: 1px solid gray;">コメント一覧</h2>

    <div
      v-for="(data, i) in this.commentsPageList"
      :key="i"
    >
      <ShowComment :commentData="data" />
    </div>

  </div>
</template>


<script>
import ShowComment from '~/components/comment/ShowComment.vue'
import { mapState, mapActions } from 'vuex'

export default {
  components: {
    ShowComment,
  },
  methods: {
    ...mapActions('comment', [
      'fetchCommentsPageList'
    ])
  },
  computed: {
    ...mapState('comment', [
      'commentsPageList'
    ])
  },
  async created() {
    await this.fetchCommentsPageList([this.$route.params.cvid, this.$route.params.page])
  }
}
</script>