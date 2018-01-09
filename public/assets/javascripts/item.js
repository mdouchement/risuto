Vue.component('item', {
  template: `
    <bulma-accordion-item>
      <div slot="title">
        {{item.name}} ・ {{item.score}}
      </div>
      <div slot="content">
        <div class="level">
          <div class="level-left">
            <ul>
              <li v-for="descrition of item.descriptions">{{descrition}}</li>
            </ul>
          </div>
          <div class="level-right">
            <div class="field is-grouped">
              <div class="secondary hollow label">{{item.score}}・</div>
              <div class="button is-primary is-outlined is-small" v-on:click="dec()">-</div>
              <div class="button is-outlined is-small" v-on:click="inc()">+</div>
            </div>
          </div>
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
      <div slot="footer">
        POPO
      </div>
    </bulma-accordion-item>
  `,
  props: ['item'],
  data: function() {
    return {
      deleteMessages: ["Delete", "Sure?"],
      deleteMessageIndex: 0,
      tmpDescription: ''
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
