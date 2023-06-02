package test

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
	"testing"
)

var resources = []model.Resource{
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
	model.NewResource("1", "kind"),
}

func insertDataCheckPermission(aclStore acl.Store) error {
	resReqs := make([]acl.CreateResourceReq, 0)

	gps := model.NewResource("1", "gps")
	ps := model.NewResource("1", "ps")
	s := model.NewResource("1", "s")
	gpo := model.NewResource("1", "gpo")
	po := model.NewResource("1", "po")
	o := model.NewResource("1", "o")
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: gps, Callback: nil})
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: ps, Callback: nil})
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: s, Callback: nil})
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: gpo, Callback: nil})
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: po, Callback: nil})
	resReqs = append(resReqs, acl.CreateResourceReq{Resource: o, Callback: nil})

	relReqs := make([]acl.CreateAggregationRelReq, 0)

	gpsPsRelReq := acl.CreateAggregationRelReq{Parent: gps, Child: ps, Callback: nil}
	psSRelReq := acl.CreateAggregationRelReq{Parent: ps, Child: s, Callback: nil}
	gpoPoRelReq := acl.CreateAggregationRelReq{Parent: gpo, Child: po, Callback: nil}
	poORelReq := acl.CreateAggregationRelReq{Parent: po, Child: o, Callback: nil}

	relReqs = append(relReqs, gpsPsRelReq)
	relReqs = append(relReqs, psSRelReq)
	relReqs = append(relReqs, gpoPoRelReq)
	relReqs = append(relReqs, poORelReq)

	permReqs := make([]acl.CreatePermissionReq, 0)
	c1, err := model.NewCondition("")
	if err != nil {
		return err
	}
	perms := make([]model.Permission, 0)
	p1 := model.NewPermission("p", model.PermissionKindAllow, *c1)
	p2 := model.NewPermission("p2", model.PermissionKindAllow, *c1)
	p3 := model.NewPermission("p3", model.PermissionKindAllow, *c1)
	perms = append(perms, p1)
	perms = append(perms, p2)
	perms = append(perms, p3)
	for _, p := range perms {
		permReq1 := acl.CreatePermissionReq{Permission: p, Subject: gps, Object: gpo, Callback: nil}
		permReq2 := acl.CreatePermissionReq{Permission: p, Subject: gps, Object: po, Callback: nil}
		permReq3 := acl.CreatePermissionReq{Permission: p, Subject: gps, Object: o, Callback: nil}
		permReq4 := acl.CreatePermissionReq{Permission: p, Subject: ps, Object: gpo, Callback: nil}
		permReq5 := acl.CreatePermissionReq{Permission: p, Subject: ps, Object: po, Callback: nil}
		permReq6 := acl.CreatePermissionReq{Permission: p, Subject: ps, Object: o, Callback: nil}
		permReq7 := acl.CreatePermissionReq{Permission: p, Subject: s, Object: gpo, Callback: nil}
		permReq8 := acl.CreatePermissionReq{Permission: p, Subject: s, Object: po, Callback: nil}
		permReq9 := acl.CreatePermissionReq{Permission: p, Subject: s, Object: o, Callback: nil}
		permReqs = append(permReqs, permReq1)
		permReqs = append(permReqs, permReq2)
		permReqs = append(permReqs, permReq3)
		permReqs = append(permReqs, permReq4)
		permReqs = append(permReqs, permReq5)
		permReqs = append(permReqs, permReq6)
		permReqs = append(permReqs, permReq7)
		permReqs = append(permReqs, permReq8)
		permReqs = append(permReqs, permReq9)
	}

	for _, req := range resReqs {
		resp := aclStore.CreateResource(req)
		if resp.Error != nil {
			return resp.Error
		}
	}
	for _, req := range relReqs {
		resp := aclStore.CreateAggregationRel(req)
		if resp.Error != nil {
			return resp.Error
		}
	}

	for _, req := range permReqs {
		resp := aclStore.CreatePermission(req)
		if resp.Error != nil {
			return resp.Error
		}
	}

	return nil
}

var subs = []model.Resource{
	model.NewResource("1", "s"),
	model.NewResource("1", "ps"),
	model.NewResource("1", "gps"),
}

var objs = []model.Resource{
	model.NewResource("1", "o"),
	model.NewResource("1", "po"),
	model.NewResource("1", "gpo"),
}

func BenchmarkCheckPermissionNoCaching(b *testing.B) {
	err := setUpAclStoreNoCaching(&AclStore, TxManager)
	if err != nil {
		b.Error(err)
	}
	err = insertDataCheckPermission(AclStore)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for _, sub := range subs {
		for _, obj := range objs {
			b.Run(fmt.Sprintf("perm sub - %s obj - %s", sub.Name(), obj.Name()), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StartTimer()
					resp := AclStore.GetPermissionHierarchy(acl.GetPermissionHierarchyReq{Subject: sub, Object: obj, PermissionName: "p"})
					b.StopTimer()
					if resp.Error != nil {
						b.Error(resp.Error)
					}
				}
			})
		}
	}
	err = cleanUpAclStore(TxManager)
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkCheckPermissionCaching(b *testing.B) {
	err := setUpAclStoreCaching(&AclStore, TxManager)
	if err != nil {
		b.Error(err)
	}
	err = insertDataCheckPermission(AclStore)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for _, sub := range subs {
		for _, obj := range objs {
			b.Run(fmt.Sprintf("perm sub - %s obj - %s", sub.Name(), obj.Name()), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StartTimer()
					resp := AclStore.GetPermissionHierarchy(acl.GetPermissionHierarchyReq{Subject: sub, Object: obj, PermissionName: "p"})
					b.StopTimer()
					if resp.Error != nil {
						b.Error(resp.Error)
					}
				}
			})
		}
	}
	err = cleanUpAclStore(TxManager)
	if err != nil {
		b.Error(err)
	}
}
