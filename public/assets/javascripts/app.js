// Vue.config.devtools = true

// Global variable
Vue.use(Buefy.default, { defaultIconPack: 'fa' })

// App
let app = new Vue({
  el: '#app',
  delimiters: ['${', '}'],
  mounted: function() {
    this.getItems()
  },
  data: {
    inNewItem: false,
    filter: '',
    items: []
  },
  methods: {
    clearFilter: function() {
      this.filter = ''
    },
    newItem: function() {
      this.inNewItem = true
    },
    closeNewItem: function() {
      this.inNewItem = false
    },
    appendItem: function(item) {
      this.items.push(item)
    },
    removeItem: function(item) {
      let i = _.findIndex(this.items, item)
      this.items.splice(i, 1)
    },
    getItems: function() {
      let self = this
      axios.get('/items', {
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        }
      })
      .then(function(response) {
        self.items = _.orderBy(response.data,['score', 'name'], ['desc', 'asc'])
      })
      .catch(function(error) {
        console.log(error)
      })
    }
  },
  computed: {
    sortedItems: function() {
      let filter = _.toLower(this.filter)
      return _.filter(this.items, i => _.startsWith(_.toLower(i.name), filter))
    }
  }
})
