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
    items: {}
  },
  methods: {
    appendCategory: function(category) {
      if (!_.find(this.categories, c => c === category)) {
        this.categories.push(category)
      }
    },
    removeCategory: function(category) {
      if (this.items[category].length === 0) {
        i = _.findIndex(this.categories, category)
        this.categories.splice(i, 1)
        // FIXME generates glitches in UI
        this.activeTab = 0
      }
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
      if (this.items[item.category] === undefined) {
        this.$set(this.items, item.category, [])
      }
      this.items[item.category].push(item)

      if (autoSwitchTab) {
        let self = this
        _.delay(item => {
          // FIXME generates glitches in UI on new fresh category
          self.activeTab = _.sortedIndex(self.sortedCategories, item.category) // Auto-switch tab
        }, 200, item)
      }
    },
    removeItem: function(item) {
      let i = _.findIndex(this.items, item)
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
        _.each(response.data, i => self.appendItem(i, false))
      })
      .catch(function(error) {
        console.log(error)
        alert(error)
      })
    }
  },
  computed: {
    sortedCategories: function() {
      return this.categories.sort()
    },
    sortedItems: function() {
      return _.orderBy(this.items[this.sortedCategories[this.activeTab]], ['score', 'name'], ['desc', 'asc'])
    },
    filteredItems: function() {
      bus.$emit('item-selected', 'collapse-all') // Force collapse all on search
      let filter = _.toLower(this.filter)
      return _.filter(this.sortedItems, i => _.startsWith(_.toLower(i.name), filter))
    }
  }
})
