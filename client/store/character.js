export const state = () => ({
  // get
  characterUnderVols: [],
  characterId: [],
  // post
  addedAppearanceCharacter: {},
  addedCharacter: {},
})

export const mutations = {
  SET_CHARACTER_UNDER_VOLS(state, list) {
    state.characterUnderVols = list
  },
  SET_CHARACTER_ID(state, list) {
    state.characterId = list
  },
  SET_ADDED_APPEARANCE_CHARACTER(state, obj) {
    state.addedAppearanceCharacter = obj
  },
  SET_ADDED_CHARACTER(state, obj) {
    state.addedCharacter = obj
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
  // nameを指定してキャラクタのidを取得
  async fetchCharacterId({ commit }, name) {
    await this.$axios({
      method: 'GET',
      url: `/characters/name/${name}`,
    }).then((res) => {
      commit('SET_CHARACTER_ID', res.data)
    }).catch(() => {})
  },
  // appearance_charactersの追加
  async postAppearanceCharacter({ commit, dispatch, rootState }, obj) {
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: '/characters/appearance',
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_APPEARANCE_CHARACTER', res.data)
    }).catch((e) => {
      console.log(e)
      statusCode = e.response.status
    })

    if (statusCode == 401) {
      await dispatch('postAppearanceCharacter', obj)
    }
  },

  // characterの追加
  async postCharacter({ commit, dispatch, rootState }, obj) {
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: `/characters`,
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_CHARACTER', res.data)
    }).catch((e) => {
      console.log(e)
      statusCode = e.response.status
    })

    if (statusCode == 401) {
      await dispatch('postCharacter', obj)
    }
  },

  // リセット
  async resetCharacter({ commit }) {
    commit('RESET_CHARACTER', [])
  }
}