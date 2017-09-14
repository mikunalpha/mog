<template>
  <div id="overview-view">
    Blog Posts Amount: {{ this.status.blog.posts.amount }}
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters({
      status: 'status'
    })
  },

  created () {
    console.log(this.status)
  }
}
</script>

<style lang="sass">
#overview-view
</style>
