Vue.component('new-item', {
  template: `
    <div class="container card x-item" v-if="active">
      <div class="card-content">
        <div class="field is-horizontal">
          <div class="field-label">
            <label class="label">Name</label>
          </div>
          <div class="field-body is-expanded">
            <input v-model="name" class="input" type="text" placeholder="Shingeki no Kyojin">
          </div>
        </div>
        <div class="field is-horizontal">
          <div class="field-label">
            <label class="label">Description</label>
          </div>
          <div class="field-body is-expanded">
            <input v-model="description" class="input" type="text" placeholder="A great anime!">
          </div>
        </div>
        <div class="level">
          <div class="level-left"></div>
          <div class="level-right">
            <div class="field is-grouped">
              <a class="button is-small" v-on:click="goBack()">Back</a>
              <a class="button is-small is-primary" v-on:click="create()">Save</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  `,
  props: {
    active: {
      type: Boolean,
      default: false
    }
  },
  data: function() {
    return {
      name: '',
      description: ''
    }
  },
  methods: {
    goBack: function() {
      this.name = ''
      this.description = ''

      this.$emit('closed')
    },
    create: function() {
      let self = this
      let descriptions = []
      if (this.description !== '') {
        descriptions = [this.description]
      }

      axios.post('/items', {
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        },
        // Data
        name: self.name,
        descriptions: descriptions
      })
      .then(function(response) {
        self.$emit('created', response.data)
      })
      .catch(function(error) {
        alert(error)
      })
      .finally(function() {
        self.goBack()
      })
    }
  }
})
