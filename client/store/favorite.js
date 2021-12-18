export const state = () => ({
  // get
  whetherFavorite: [],
  favoritesList: [],

  // post
  addedFavoriteData: {},

})

export const mutations = {
  SET_WHETHER_FAVORITE(state, list) {
    state.whetherFavorite = list
  },
  SET_FAVORITES_LIST(state, list) {
    state.favoritesList = list
  },
  SET_ADDED_FAVORITE_DATA(state, obj) {
    state.addedFavoriteData = obj
  },
  RESET_FAVORITE_DATA(state, n) {
    state.whetherFavorite = n
  }
}

export const actions = {
  // uidとscidを指定してお気に入りかどうか判断
  async fetchWhetherFavorite({ commit }, [uid, scid]) {
    await this.$axios({
      method: 'GET',
      url: `/favorites/whether/${uid}/${scid}`,
    }).then((res) => {
      commit('SET_WHETHER_FAVORITE', res.data)
    }).catch(() => {})
  },
  // uidを指定してユーザのお気に入りシーン一覧を取得
  async fetchFavoritesList({ commit }, uid) {
    await this.$axios({
      method: 'GET',
      url: `/favorites/list/${uid}`,
    }).then((res) => {
      commit('SET_FAVORITES_LIST', res.data)
    }).catch(() => {})
  },

  // post
  // お気に入りのついか
  async postFavorite({ commit, dispatch, rootState }, obj) {
    console.log("Q")
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: '/favorites',
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_FAVORITE_DATA', res.data)
    }).catch((e) => {
      statusCode = e.response.status
    })

    if (statusCode == 401) {
      await dispatch('postFavorite', obj)
    }
  },

  //delete お気に入りの削除
  async deleteFavorite({ commit }, [uid, scid]) {
    await this.$axios({
      method: 'DELETE',
      url: `/favorites/${uid}/${scid}`
    }).then((res) => {
      console.log('削除しました')
    })
  },

  //リセット
  async resetFavorite({ commit }) {
    commit('RESET_FAVORITE_DATA', [])
  }
}