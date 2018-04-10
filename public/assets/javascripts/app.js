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
    categories: [],
    activeTab: 0,
    items: {},
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
    appendItem: function(item, pos=-1, autoSwitchTab=true) {
      this.appendCategory(item.category)
      if (this.items[item.category] === undefined) {
        this.$set(this.items, item.category, [])
      }
      this.items[item.category].push(item)


      if (autoSwitchTab) {
        let self = this
        _.delay(item => {
          // FIXME generates glitches in UI on new fresh category
          self.activeTab = _.indexOf(self.categories, item.category) // Auto-switch tab
        }, 200, item)
      }
    },
    removeItem: function(item) {
      let i = _.findIndex(this.items[item.category], item)
      this.items[item.category].splice(i, 1)
      this.removeCategory(item.category)
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
        let items = _.orderBy(response.data, ['score', 'name'], ['desc', 'asc'])
        _.each(items, i => {
          self.appendItem(i, false)
        })
        self.categories = self.categories.sort()
      })
      .catch(function(error) {
        console.log(error)
        alert(error)
      })
    }
  },
  computed: {
    itemPool: function() {
      return this.items[this.categories[this.activeTab]]
    },
    filteredItems: function() {
      bus.$emit('item-selected', 'collapse-all') // Force collapse all on search
      let filter = _.toLower(this.filter)
      return _.filter(this.itemPool, i => _.startsWith(_.toLower(i.name), filter))
    }
  }
})
