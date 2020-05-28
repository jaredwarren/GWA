package ext

func buildTree(i interface{}) *Tree {
	ii := i.(map[string]interface{})

	p := &Tree{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if showRoot, ok := ii["showRoot"]; ok {
		p.ShowRoot = showRoot.(bool)
	}

	if branchIcon, ok := ii["branchIcon"]; ok {
		p.BranchIcon = branchIcon.(string)
	}

	if leafIcon, ok := ii["leafIcon"]; ok {
		p.LeafIcon = leafIcon.(string)
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	if root, ok := ii["root"]; ok {
		p.Root = buildTreeNode(root)
	}

	return p
}

func buildTreeNode(i interface{}) *TreeNode {
	ii := i.(map[string]interface{})

	p := &TreeNode{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if collapsed, ok := ii["collapsed"]; ok {
		p.Collapsed = collapsed.(bool)
	}

	if leaf, ok := ii["leaf"]; ok {
		p.Leaf = leaf.(bool)
	}

	if text, ok := ii["text"]; ok {
		p.Text = text.(string)
	}

	if iconClass, ok := ii["iconClass"]; ok {
		p.IconClass = iconClass.(string)
	}

	items := []*TreeNode{}
	if ii, ok := ii["children"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := buildTreeNode(i)
			items = append(items, item)
		}
	}
	p.Children = items

	return p
}
