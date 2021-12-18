import { getAuth, signInWithPopup, GoogleAuthProvider } from 'firebase/auth'


export const state = () => ({
  userData: {},
  jwt: null,
})

export const getters = {
  isLoggedIn(state) {
    return !!Object.keys(state.userData).length
  },
}

export const mutations = {
  SET_USER_DATA(state, data) {
    state.userData = data
  },
  SET_JWT(state, token) {
    state.jwt = token
  },
}

export const actions = {
  // Firebase関連
  firebaseAuth() {
    const provider = new GoogleAuthProvider()
    const auth = getAuth()

    signInWithPopup(auth, provider).then((res) => {
      res.user.getIdToken().then((idToken) => {
        localStorage.setItem('jwt', idToken.toString())
      })
    })
  },
  async refreshToken() {
    const auth = getAuth()
    const user = auth.currentUser
    if (user) {
      await user.getIdToken(true)
    }
  },
  async signIn({ commit, dispatch, rootState }, user) {
    let isExist = false
    await this.$axios({
      method: 'GET',
      url: '/users/me',
      params: {
        Mail: user.email,
      },
    })
      .then((res) => {
        isExist = true
        commit('SET_USER_DATA', res.data)
      })
      .catch((e) => {
      })

    if (!isExist) {
      await this.$axios({
        method: 'POST',
        url: '/users',
        headers: { Authorization: `Bearer ${rootState.user.jwt}` },
        data: {
          Name: user.displayName,
          Mail: user.email,
          Image: user.photoURL,
        },
      }).then((res) => {
        commit('SET_USER_DATA', res.data)
      })
    }
  },
  signOut({ commit }) {
    const auth = getAuth()
    auth.signOut()
    commit('SET_USER_DATA', {})
    commit('SET_JWT', null)
  },
  setJwt({ commit }, token) {
    commit('SET_JWT', token)
  },
  setUserData({ commit }, user) {
    commit('SET_USER_DATA', user)
  },
}