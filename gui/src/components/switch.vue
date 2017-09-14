<template>
  <div class="switch">
    <template v-if="options.length > 0">
      <div class="switch-options">
        <div v-for="(option, index) in options"
          class="switch-option"
          :class="{selected: (index == selectedOptionIndex)}"
          @click="selectOption(index)">
          {{ option.text }}
        </div>
      </div>  
    </template>
  </div>
</template>

<script>
export default {
  props: {
    options: {
      type: Array,
      default () { return [] }
    }
  },

  data () {
    return {
      selectedOptionIndex: 0
    }
  },

  methods: {
    selectOption (i) {
      this.selectedOptionIndex = i
      this.$emit('select', {
        index: i,
        option: this.options[i]
      })
    }
  },

  created () {
    for (let i in this.options) {
      if (this.options[i].selected === true) {
        this.selectOption(i)
        break
      }
    }
  }
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

.switch
  display: inline-block

  .switch-options
    display: flex
    padding: 5px
    box-shadow: 1px 1px 0 darken($backgroundColor, 20%) inset, -1px -1px 0 darken($backgroundColor, 20%) inset

    .switch-option
      flex: 1
      display: flex
      align-items: center
      text-align: center
      padding: 5px 8px
      cursor: pointer
      font-size: 12px
      font-weight: bold
      &.selected
        background-color: darken($backgroundColor, 50%)
        color: #ffffff
</style>
