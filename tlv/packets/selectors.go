package packets

import "ndn-router/nfd/tlv/name"

// Selectors ::= SELECTORS-TYPE TLV-LENGTH
//                 MinSuffixComponents?
//                 MaxSuffixComponents?
//                 PublisherPublicKeyLocator?
//                 Exclude?
//                 ChildSelector?
//                 MustBeFresh?
type Selectors struct {
	HasMinSuffixComponents       bool
	minSuffixComponents          uint64
	HasMaxSuffixComponents       bool
	maxSuffixComponents          uint64
	HasPublisherPublicKeyLocator bool
	publisherPublicKeyLocator    KeyLocator
	HasExclude                   bool
	exclude                      name.Exclude
	HasChildSelector             bool
	childSelector                uint64
	mustBeFresh                  bool
}

func (sel Selectors) IsEmpty() bool {
	return !(sel.HasMinSuffixComponents && sel.HasMaxSuffixComponents && sel.HasPublisherPublicKeyLocator && sel.HasChildSelector && sel.HasExclude && sel.mustBeFresh)
}

func (sel Selectors) GetMinSuffixComponents() uint64 {
	if !sel.HasMinSuffixComponents {
		return 0
	}
	return sel.minSuffixComponents
}

func (sel *Selectors) SetMinSuffixComponents(x uint64) {
	sel.HasMinSuffixComponents = true
	sel.minSuffixComponents = x
}

func (sel Selectors) GetMaxSuffixComponents() uint64 {
	if !sel.HasMaxSuffixComponents {
		return 0
	}
	return sel.maxSuffixComponents
}

func (sel *Selectors) SetMaxSuffixComponents(x uint64) {
	sel.HasMaxSuffixComponents = true
	sel.maxSuffixComponents = x
}

func (sel Selectors) GetPublisherPublicKeyLocator() KeyLocator {
	if !sel.HasPublisherPublicKeyLocator {
		return KeyLocator{}
	}
	return sel.publisherPublicKeyLocator
}

func (sel *Selectors) SetPublisherPublicKeyLocator(pKey KeyLocator) {
	sel.HasPublisherPublicKeyLocator = true
	sel.publisherPublicKeyLocator = pKey
}

func (sel Selectors) GetExclude() name.Exclude {
	if !sel.HasExclude {
		return nil
	}
	return sel.exclude
}

func (sel *Selectors) SetExclude(ex name.Exclude) {
	sel.HasExclude = true
	sel.exclude = ex
}

func (sel Selectors) GetChildSelector() uint64 {
	if !sel.HasChildSelector {
		return 0
	}
	return sel.childSelector
}

func (sel *Selectors) SetChildSelector(cs uint64) {
	sel.HasChildSelector = true
	sel.childSelector = cs
}

func (sel Selectors) GetMustBeFresh() bool {
	return sel.mustBeFresh
}

func (sel *Selectors) SetMustBeFresh(mbf bool) {
	sel.mustBeFresh = mbf
}
