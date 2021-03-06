package configs

func findSource(user []string, source string) (int) {
	index := -1
	for i, s := range(user) {
		if s == source {
			index = i
			break
		}
	}
	return index
}

func (c *Configs) UserAddSource(uid int, source string) {
	c.ulock.RLock()
	u, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	if !ok {
		u = make([]string, 1)
		u[0] = source
		c.ulock.Lock()
		c.contents.Users[uid] = u
		c.ulock.Unlock()
	} else {
		if findSource(u, source) < 0 {
			c.ulock.Lock()
			c.contents.Users[uid] = append(u, source)
			c.ulock.Unlock()
		}
	}
	c.Store()
}

func (c *Configs) UserDelSource(uid int, source string) {
	c.ulock.RLock()
	u, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	if ok {
		index := findSource(u, source)
		if index >= 0 {
			c.contents.Users[uid] = append(u[:index], u[index+1:]...)
		}
		c.Store()
	}
}

func (c *Configs) UserToggleSource(uid int, source string) (added bool) {
	c.ulock.RLock()
	u, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	if ok {
		index := findSource(u, source)
		if index < 0 {
			c.ulock.Lock()
			c.contents.Users[uid] = append(u, source)
			c.ulock.Unlock()
			added = true
		} else {
			c.ulock.Lock()
			c.contents.Users[uid] = append(u[:index], u[index+1:]...)
			c.ulock.Unlock()
			added = false
		}
	} else {
		c.UserAddSource(uid, source)
		added = true
	}
	c.Store()
	return added
}

func (c *Configs) UserSourceEnabled(uid int, source string) (enabled bool) {
	c.ulock.RLock()
	u, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	if ok {
		if findSource(u, source) < 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func (c *Configs) GetUser(uid int) ([]string, bool) {
	c.ulock.RLock()
	u, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	return u, ok
}

func (c *Configs) GetUsers(source string) ([]int) {
	c.ulock.RLock()
	uids := make([]int, 0, len(c.contents.Users))
	for u, list := range c.contents.Users {
		for _, s := range list {
			if source == s {
				uids = append(uids, u)
				break
			}
		}
	}
	c.ulock.RUnlock()
	return uids
}

func (c *Configs) GetAllUsers() ([]int) {
	c.ulock.RLock()
	uids := make([]int, len(c.contents.Users))
	i := 0
	for u := range c.contents.Users{
		uids[i] = u
		i++
	}
	c.ulock.RUnlock()
	return uids
}

func (c *Configs) CreateUser(uid int) {
	c.ulock.RLock()
	_, ok := c.contents.Users[uid]
	c.ulock.RUnlock()
	if !ok {
		c.ulock.Lock()
		c.contents.Users[uid] = make([]string, 0)
		c.ulock.Unlock()
		c.Store()
	}
}
