package ext

import "html/template"

func buildForm(i interface{}) *Form {
	ii := i.(map[string]interface{})

	p := &Form{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}

	if action, ok := ii["action"]; ok {
		p.Action = action.(string)
	}

	if method, ok := ii["method"]; ok {
		p.Method = method.(string)
	}

	if handler, ok := ii["handler"]; ok {
		p.Handler = handler.(string)
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

	items := Items{}
	if ii, ok := ii["items"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := addChild(i)
			items = append(items, item)
		}
	}

	p.Items = items

	return p
}

func buildFieldset(i interface{}) *Fieldset {
	ii := i.(map[string]interface{})

	p := &Fieldset{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
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

	items := []Renderer{}
	if ii, ok := ii["items"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := addChild(i)
			items = append(items, item)
		}
	}

	if legend, ok := ii["legend"]; ok {
		p.Legend = template.HTML(legend.(string))
	}

	p.Items = items
	return p
}

func buildInput(i interface{}) *Input {
	ii := i.(map[string]interface{})

	p := &Input{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if t, ok := ii["type"]; ok {
		p.Type = t.(string)
	}

	if name, ok := ii["name"]; ok {
		p.Name = name.(string)
	}

	if value, ok := ii["value"]; ok {
		p.Value = value.(string)
	}

	if form, ok := ii["form"]; ok {
		p.Form = form.(string)
	}

	if disabled, ok := ii["disabled"]; ok {
		p.Disabled = disabled.(bool)
	}

	if autofocus, ok := ii["autofocus"]; ok {
		p.Autofocus = autofocus.(bool)
	}

	if autocomplete, ok := ii["autocomplete"]; ok {
		p.Autocomplete = autocomplete.(string)
	}

	if label, ok := ii["label"]; ok {
		p.Label = template.HTML(label.(string))
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

	if s, ok := ii["attributes"]; ok {
		jclass := s.(map[string]interface{})
		attributes := map[string]template.HTMLAttr{}
		for i, cl := range jclass {
			attributes[i] = template.HTMLAttr(cl.(string))
		}
		p.Attributes = attributes
	}

	return p
}
