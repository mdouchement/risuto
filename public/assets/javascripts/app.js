// Vue.config.devtools = true

// Global variable
Vue.use(Buefy.default, { defaultIconPack: 'fa' })

// App
let app = new Vue({
  el: '#app',
  delimiters: ['${', '}'],
  mounted: function() {
    this.getCategories()
  },
  data: {
    inNewItem: false,
    filter: '',
    categories: [],
    activeTab: 0,
    items: [],
  },
  methods: {
    appendCategory: function(category) {
      if (!_.find(this.categories, c => c === category)) {
        this.categories.push(category)
      }
    },
    removeCategory: function(category) {
      // if (this.items_map[category].length === 0) {
      //   i = _.findIndex(this.categories, category)
      //   this.categories.splice(i, 1)
      //   // FIXME generates glitches in UI
      //   this.activeTab = 0
      // }
    },
    getCategories: function() {
      let self = this
      axios.get('/categories', {
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        }
      })
      .then(response => {
        self.categories = response.data.sort()
        self.getItems()
      })
      .catch(error => {
        console.log(error)
        alert(error)
      })
    },
    clearFilter: function() {
      this.filter = ''
    },
    newItem: function() {
      bus.$emit('item-selected', 'collapse-all')
      this.inNewItem = true
    },
    closeNewItem: function() {
      this.inNewItem = false
    },
    appendItem: function(item, autoSwitchTab=true) {
      this.appendCategory(item.category)
      this.items.push(item)


      if (autoSwitchTab) {
        let self = this
        _.delay(item => {
          self.activeTab = _.indexOf(self.categories, item.category) // Auto-switch tab
        }, 200, item)
      }
    },
    removeItem: function(item) {
      let i = _.findIndex(this.items, item)
      this.items.splice(i, 1)
      this.removeCategory(item.category)
    },
    refreshItems: function() {
      this.items = []
      this.getItems()
    },
    getItems: function() {
      let self = this
      axios.get('/items?category='+self.categories[self.activeTab], {
        headers: {
          'Content-type': 'application/json',
          'Accept': 'application/json'
        }
      })
      .then(response => {
        let items = _.orderBy(response.data, ['score', 'name'], ['desc', 'asc'])
        _.each(items, i => {
          self.appendItem(i, false)
        })
      })
      .catch(error => {
        console.log(error)
        alert(error)
      })
    }
  },
  computed: {
    filterDebounced: {
      get: function() {
        return this.filter
      },
      set: _.debounce(function(filter) {
        this.filter = filter
      }, 500)
    },
    filteredItems: function() {
      bus.$emit('item-selected', 'collapse-all') // Force collapse all on search
      if (this.filter === "") {
        return this.items
      }
      let filter = _.toLower(this.filter)
      return _.filter(this.items, i => i.name.toLowerCase().includes(filter))
    }
  }
})
