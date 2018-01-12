let bus = new Vue()

Vue.component('item', {
  template: `
    <b-collapse class="card" v-bind:class="{ 'x-item': isOpen }" :open.sync="isOpen">
      <div slot="trigger" class="card-header">
        <p class="card-header-title">
          {{item.name}} ・ {{item.score}}
        </p>
        <a class="card-header-icon">
          <b-icon v-bind:icon="isOpen ? 'angle-down' : 'angle-up'"></b-icon>
        </a>
      </div>
      <div class="card-content">
        <div class="content">
          <div class="level">
            <div class="level-left">
              <button v-bind:id="copyButtonID" class="button is-small" v-bind:data-clipboard-text="item.name">
                <b-icon icon="clipboard"></b-icon>
              </button>
            </div>
            <div class="level-right">
              <div class="field is-grouped">
                <div class="secondary hollow label">{{item.score}}・</div>
                <div class="button is-primary is-outlined is-small" v-on:click="dec()">-</div>
                <div class="button is-outlined is-small" v-on:click="inc()">+</div>
              </div>
            </div>
          </div>

          <div class="field">
            <ul>
              <li v-for="descrition of item.descriptions">{{descrition}}</li>
            </ul>
          </div>

          <div class="field is-grouped">
            <p class="control is-expanded">
              <input class="input" type="text" v-model="tmpDescription" placeholder="Add description">
            </p>
            <p class="control">
              <a class="button is-info" v-on:click="add()">+</a>
            </p>
          </div>

          <div class="level">
            <div class="level-left"></div>
            <div class="level-right">
              <a class="level-item button is-danger is-small is-outlined" v-on:click="destroy()">{{deleteMessages[deleteMessageIndex]}}</a>
            </div>
          </div>
        </div>
      </div>
    </b-collapse>
  `,
  props: ['item'],
  data: function() {
    return {
      isOpen: false,
      deleteMessages: ["Delete", "Sure?"],
      deleteMessageIndex: 0,
      tmpDescription: '',
      clipboard: null
    }
  },
  mounted() {
    let self = this
    bus.$on('item-selected', function(id) {
      self.isOpen = self.item.id === id
    })
  },
  beforeDestroy() {
    if (this.clipboard === null) {
      this.clipboard.destroy()
    }
    bus.$emit('item-selected', 'collapse-all')
  },
  watch: {
    isOpen: function(active) {
      if (this.clipboard === null) {
        this.clipboard = new Clipboard('#' + this.copyButtonID) // Only on-demand
      }

      if (active) {
        bus.$emit('item-selected', this.item.id)
      }
    },
    deleteMessageIndex: function(index) {
      if (index > 0) {
        let self = this
        _.delay(function() { self.deleteMessageIndex = 0 }, 1000)
      }
    }
  },
  computed: {
    copyButtonID: function() {
      return 'item-button-' + this.item.id
    }
  },
  methods: {
    add: function() {
      this.item.descriptions.push(this.tmpDescription)
      this.tmpDescription = ''
      this.update()
    },
    inc: function() {
      this.item.score++
      this.update()
    },
    dec: function() {
      if (this.item.score > 0) {
        this.item.score--
        this.update()
      }
    },
    update: function() {
      let self = this
      axios.patch('/items/'+self.item.id, _.merge({
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        },
      }, self.item))
      .catch(function(error) {
        alert(error)
      })
    },
    destroy: function() {
      if (this.deleteMessageIndex < 1) {
        this.deleteMessageIndex++
        return
      }

      let self = this
      axios.delete('/items/'+self.item.id, {
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        },
      })
      .catch(function(error) {
        alert(error)
      })
      .finally(function() {
        self.$emit('deleted', self.item)
      })
    }
  }
})
