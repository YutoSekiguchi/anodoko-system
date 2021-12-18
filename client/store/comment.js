export const state = () => ({
  // get
  commentsPageList: [],
})

export const mutations = {
  SET_COMMENTS_PAGE_LIST(state, list) {
    state.commentsPageList = list
  },
}

export const actions = {
  // cvidとpageを指定してコメントを取得
  async fetchCommentsPageList({ commit }, [cvid, page]) {
    await this.$axios({
      method: 'GET',
      url: `/comments/list/${cvid}/${page}`,
    }).then((res) => {
      commit('SET_COMMENTS_PAGE_LIST', res.data)
    }).catch(() => {})
  },
}