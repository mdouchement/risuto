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
    items: [],
    items_map: {}
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
      if (this.items_map[item.category] === undefined) {
        this.$set(this.items_map, item.category, [])
      }
      if (pos == -1) {
        // New item
        this.items_map[item.category].push(this.items.length)
        this.items.push(item)
      } else {
        // Already added item
        this.items_map[item.category].push(pos)
      }

      if (autoSwitchTab) {
        let self = this
        _.delay(item => {
          // FIXME generates glitches in UI on new fresh category
          self.activeTab = _.indexOf(self.categories, item.category) // Auto-switch tab
        }, 200, item)
      }
    },
    removeItem: function(item) {
      let pos = _.findIndex(this.items, item)
      let i = _.findIndex(this.items_map[item.category], pos)
      this.items_map[item.category].splice(i, 1)
      // Sync mapping
      _.each(this.categories, c => {
        this.items_map[c] = _.map(this.items_map[c], cpos => cpos > pos ? cpos-1 : cpos)
      })
      this.items.splice(pos, 1)
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
        self.items = _.orderBy(response.data, ['score', 'name'], ['desc', 'asc'])
        let pos = 0
        _.each(self.items, i => {
          self.appendItem(i, pos, false)
          pos++
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
      return _.map(this.items_map[this.categories[this.activeTab]], pos => this.items[pos])
    },
    filteredItems: function() {
      bus.$emit('item-selected', 'collapse-all') // Force collapse all on search
      let filter = _.toLower(this.filter)
      return _.filter(this.itemPool, i => _.startsWith(_.toLower(i.name), filter))
    }
  }
})
