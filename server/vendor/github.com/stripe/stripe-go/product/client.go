package product

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /products APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func New(params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().New(params)
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func (c Client) New(params *stripe.ProductParams) (*stripe.Product, error) {
	p := &stripe.Product{}
	err := c.B.Call(http.MethodPost, "/products", c.Key, params, p)
	return p, err
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Update(id, params)
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func (c Client) Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/products/%s", id)
	p := &stripe.Product{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, p)
	return p, err
}

// Get returns the details of an product
// For more details see https://stripe.com/docs/api#retrieve_product.
func Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/products/%s", id)
	p := &stripe.Product{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, p)
	return p, err
}

// List returns a list of products.
// For more details see https://stripe.com/docs/api#list_products
func List(params *stripe.ProductListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.ProductListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ProductList{}
		err := c.B.CallRaw(http.MethodGet, "/products", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Products.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Product returns the most recent Product
// visited by a call to Next.
func (i *Iter) Product() *stripe.Product {
	return i.Current().(*stripe.Product)
}

// Delete deletes a product
// For more details see https://stripe.com/docs/api#delete_product.
func Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Del(id, params)
}

// Delete deletes a product.
// For more details see https://stripe.com/docs/api#delete_product.
func (c Client) Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/products/%s", id)
	p := &stripe.Product{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, p)

	return p, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
