<template>
  <div id="login-view">
    <template v-if="status && status.administered">
      <div class="title">
        Mog
      </div>

      <template v-if="errorMessage !== ''">
        <ui-alert
          class="error-message"
          type="error"
          @dismiss="clearErrorMessage">
          {{ errorMessage }}
        </ui-alert>
      </template>

      <div class="panel">
        <ui-textbox
          class="username"
          type="text"
          icon="person"
          placeholder="Username"
          v-model="credentials.username">
        </ui-textbox>

        <ui-textbox
          class="password"
          type="password"
          icon="lock"
          placeholder="Password"
          v-model="credentials.password">
        </ui-textbox>

        <br/>

        <div class="actions">
          <div class="left">

          </div>
          <div class="right">
            <ui-button
              type="primary"
              color="primary"
              :loading="isLogining"
              @click="login">
              Login
            </ui-button>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import api from '@/api/auth'

export default {
  data () {
    return {
      credentials: {
        username: '',
        password: ''
      },
      isLogining: false,
      errorMessage: ''
    }
  },

  computed: {
    ...mapGetters({
      status: 'status'
    })
  },

  methods: {
    ...mapActions({
      getAuthInfo: 'getAuthInfo'
    }),

    login () {
      let self = this

      if (self.credentials.username.trim().length === 0) {
        self.errorMessage = 'Username was required.'
        return
      } else {
        self.credentials.username = self.credentials.username.trim()
      }
      if (self.credentials.password.trim().length === 0) {
        self.errorMessage = 'Password was required.'
        return
      } else {
        self.credentials.password = self.credentials.password.trim()
      }

      self.isLogining = true

      api.login({
        credentials: self.credentials,
        successCallback: (data) => {
          self.clearErrorMessage()
          self.getAuthInfo({})
          self.$router.push({name: 'Blog.Posts'})
        },
        errorCallback: ({status, error}) => {
          self.isLogining = false
          if (status === 404) {
            self.errorMessage = 'Provided username or password was incorrect.'
          }
        }
      })
    },

    clearErrorMessage () {
      this.errorMessage = ''
    }
  }
}
</script>

<style lang="sass">
$panelWidth: 300px

#login-view
  padding-top: 100px
  min-height: 100vh
  background-color: #f0f0f0
  .title
    padding-bottom: 20px
    text-align: center
    font-size: 1.6rem
  .error-message
    margin: 0 auto
    width: $panelWidth
  .panel
    margin: 0 auto
    padding: 40px 30px 30px 30px
    width: $panelWidth
    background-color: #ffffff
    box-shadow: 0 1px 0 #e0e0e0
    .username, .password
      input
        padding-left: 5px
        letter-spacing: 4px
    .actions
      display: flex
      > div
        flex-grow: 1
      .right
        text-align: right
</style>
