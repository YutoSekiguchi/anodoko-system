export const state = () => ({
  // get
  characterUnderVols: [],

  // post
  
})

export const mutations = {
  SET_CHARACTER_UNDER_VOLS(state, list) {
    state.characterUnderVols = list
  },
  RESET_CHARACTER(state, n) {
    state.characterUnderVols = n
  }
}

export const actions = {
  // 指定した漫画の巻以下のキャラクターを取得
  async fetchCharacterUnderVols({ commit }, [cid, volume]) {
    await this.$axios({
      method: 'GET',
      url: `/characters/under/${cid}/${volume}`,
    }).then((res) => {
      commit('SET_CHARACTER_UNDER_VOLS', res.data)
    }).catch(() => {})
  },
  // リセット
  async resetCharacter({ commit }) {
    commit('RESET_CHARACTER', [])
  }
}