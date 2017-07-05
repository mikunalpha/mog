<template>
  <div id="blog-view">
    <header>
      <ui-toolbar
        type="colored"
        textColor="white"
        :raised="false"
        :removeBrandDivider="false"
        :removeNavIcon="false">
        <div slot="icon">

        </div>

        <div slot="brand">
          <span class="brand">Mog</span>
        </div>

        <div slot="actions">
          <router-view
            ref="Actions"
            name="actions"
            @channel="toDefault">
          </router-view>
        </div>
      </ui-toolbar>
    </header>

    <router-view
      ref="Default"
      name="default"
      @channel="toActions">
    </router-view>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  computed: {
    ...mapGetters({
      authInfo: 'authInfo'
    })
  },

  methods: {
    ...mapActions({
      getAuthInfo: 'getAuthInfo'
    }),
    toDefault (payload) {
      this.$refs.Default.in(payload)
    },
    toActions (payload) {
      this.$refs.Actions.in(payload)
    }
  },

  mounted () {
    this.getAuthInfo({}) // will be deprecated
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#blog-view
  position: relative
  > header
    padding: 0 20px
    background-color: $primaryColor
    .ui-toolbar
      margin: 0 auto
      max-width: 960px
      height: 60px
      .brand
        padding: 0 10px
        cursor: default
        user-select: none
        font-size: 110%
</style>
