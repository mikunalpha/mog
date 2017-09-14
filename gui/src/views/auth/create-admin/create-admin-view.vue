<template>
  <div id="new-admin-view">
    <div class="title">
      Create Admin
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
        v-model="account.username">
      </ui-textbox>

      <ui-textbox
        class="password"
        type="password"
        icon="lock"
        placeholder="Password"
        v-model="account.password">
      </ui-textbox>

      <ui-textbox
        class="password"
        type="password"
        icon="lock"
        placeholder="Confirm Password"
        v-model="confirmPassword">
      </ui-textbox>

      <br/>

      <div class="actions">
        <div class="left">

        </div>
        <div class="right">
          <ui-button
            type="primary"
            color="primary"
            :loading="isSaving"
            @click="createAdmin">
            Send
          </ui-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import api from '@/api/account'

export default {
  name: 'new-admin-view',

  data () {
    return {
      account: {
        username: '',
        password: ''
      },
      confirmPassword: '',
      isSaving: false,
      errorMessage: ''
    }
  },

  computed: {
    ...mapGetters({
      status: 'status',
      roles: 'roles'
    })
  },

  methods: {
    ...mapActions({
      getStatus: 'getStatus'
    }),

    createAdmin () {
      let self = this

      if (self.account.username.trim().length === 0) {
        self.errorMessage = 'Username was required.'
        return
      } else {
        self.account.username = self.account.username.trim()
      }
      if (self.account.password.trim().length === 0) {
        self.errorMessage = 'Password was required.'
        return
      } else {
        self.account.password = self.account.password.trim()
      }
      if (self.account.password !== self.confirmPassword) {
        return
      }

      self.isSaving = true
      self.account.role = this.roles.Admin

      api.createAccount({
        newAccount: self.account,
        successCallback: (data) => {
          self.getStatus({})
          self.$router.replace({name: 'Auth.Login'})
        },
        errorCallback: ({status, error}) => {
          self.isSaving = false
        }
      })
    },

    clearErrorMessage () {
      this.errorMessage = ''
    }
  },

  mounted () {
    if (this.status.gotAt !== null && this.status.administered) {
      this.$router.replace({name: 'Blog.Posts'})
    }
  }
}
</script>

<style lang="sass">
$panelWidth: 360px

#new-admin-view
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
