// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testclient

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"go.infratographer.com/x/gidx"
)

// An object with an ID.
// Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
type Node interface {
	IsNode()
	// The id of the object.
	GetID() gidx.PrefixedID
}

type Entity interface {
	IsEntity()
}

// Input information to create a virtual machine cpu config.
type CreateVirtualMachineCPUConfigInput struct {
	// The number of cores for this virtual machine.
	Cores int64 `json:"cores"`
	// The number of sockets for this virtual machine.
	Sockets          int64            `json:"sockets"`
	VirtualMachineID *gidx.PrefixedID `json:"virtualMachineID,omitempty"`
}

// Create a new virtual machine.
type CreateVirtualMachineInput struct {
	// The name of the Virtual Machine.
	Name string `json:"name"`
	// The ID for the owner of this Virtual Machine.
	OwnerID gidx.PrefixedID `json:"ownerID"`
	// The ID for the location of this virtual machine.
	LocationID gidx.PrefixedID `json:"locationID"`
	// The userdata for this virtual machine.
	Userdata                  *string         `json:"userdata,omitempty"`
	VirtualMachineCPUConfigID gidx.PrefixedID `json:"virtualMachineCPUConfigID"`
}

type Location struct {
	ID              gidx.PrefixedID          `json:"id"`
	VirtualMachines VirtualMachineConnection `json:"virtualMachines"`
}

func (Location) IsEntity() {}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type ResourceOwner struct {
	ID             gidx.PrefixedID          `json:"id"`
	VirtualMachine VirtualMachineConnection `json:"virtualMachine"`
}

func (ResourceOwner) IsEntity() {}

// Input information to update a virtual machine cpu config.
type UpdateVirtualMachineCPUConfigInput struct {
	// The number of cores for this virtual machine.
	Cores *int64 `json:"cores,omitempty"`
	// The number of sockets for this virtual machine.
	Sockets *int64 `json:"sockets,omitempty"`
}

// Update an existing virtual machine.
type UpdateVirtualMachineInput struct {
	// The name of the Virtual Machine.
	Name *string `json:"name,omitempty"`
	// The userdata for this virtual machine.
	Userdata      *string `json:"userdata,omitempty"`
	ClearUserdata *bool   `json:"clearUserdata,omitempty"`
}

type VirtualMachine struct {
	// The ID of the VirtualMachine.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// The name of the Virtual Machine.
	Name string `json:"name"`
	// The userdata for this virtual machine.
	Userdata *string `json:"userdata,omitempty"`
	// The ID for the virtual machine cpu config.
	VMCPUConfigID gidx.PrefixedID `json:"vmCPUConfigID"`
	// The virtual machine cpu config for the virtual machine.
	VirtualMachineCPUConfig VirtualMachineCPUConfig `json:"virtualMachineCPUConfig"`
	// The location of the load balancer.
	Location Location `json:"location"`
	// The owner of the VirtualMachine
	Owner ResourceOwner `json:"owner"`
}

func (VirtualMachine) IsNode() {}

// The id of the object.
func (this VirtualMachine) GetID() gidx.PrefixedID { return this.ID }

func (VirtualMachine) IsEntity() {}

type VirtualMachineCPUConfig struct {
	// The ID for the virtual machaine cpu config.
	ID gidx.PrefixedID `json:"id"`
	// The number of cores for this virtual machine.
	Cores int64 `json:"cores"`
	// The number of sockets for this virtual machine.
	Sockets        int64           `json:"sockets"`
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`
}

func (VirtualMachineCPUConfig) IsNode() {}

// The id of the object.
func (this VirtualMachineCPUConfig) GetID() gidx.PrefixedID { return this.ID }

func (VirtualMachineCPUConfig) IsEntity() {}

// A connection to a list of items.
type VirtualMachineCPUConfigConnection struct {
	// A list of edges.
	Edges []*VirtualMachineCPUConfigEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type VirtualMachineCPUConfigEdge struct {
	// The item at the end of the edge.
	Node *VirtualMachineCPUConfig `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for VirtualMachineCPUConfig connections
type VirtualMachineCPUConfigOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order VirtualMachineCPUConfigs.
	Field VirtualMachineCPUConfigOrderField `json:"field"`
}

// VirtualMachineCPUConfigWhereInput is used for filtering VirtualMachineCPUConfig objects.
// Input was generated by ent.
type VirtualMachineCPUConfigWhereInput struct {
	Not *VirtualMachineCPUConfigWhereInput   `json:"not,omitempty"`
	And []*VirtualMachineCPUConfigWhereInput `json:"and,omitempty"`
	Or  []*VirtualMachineCPUConfigWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// cores field predicates
	Cores      *int64  `json:"cores,omitempty"`
	CoresNeq   *int64  `json:"coresNEQ,omitempty"`
	CoresIn    []int64 `json:"coresIn,omitempty"`
	CoresNotIn []int64 `json:"coresNotIn,omitempty"`
	CoresGt    *int64  `json:"coresGT,omitempty"`
	CoresGte   *int64  `json:"coresGTE,omitempty"`
	CoresLt    *int64  `json:"coresLT,omitempty"`
	CoresLte   *int64  `json:"coresLTE,omitempty"`
	// sockets field predicates
	Sockets      *int64  `json:"sockets,omitempty"`
	SocketsNeq   *int64  `json:"socketsNEQ,omitempty"`
	SocketsIn    []int64 `json:"socketsIn,omitempty"`
	SocketsNotIn []int64 `json:"socketsNotIn,omitempty"`
	SocketsGt    *int64  `json:"socketsGT,omitempty"`
	SocketsGte   *int64  `json:"socketsGTE,omitempty"`
	SocketsLt    *int64  `json:"socketsLT,omitempty"`
	SocketsLte   *int64  `json:"socketsLTE,omitempty"`
	// virtual_machine edge predicates
	HasVirtualMachine     *bool                       `json:"hasVirtualMachine,omitempty"`
	HasVirtualMachineWith []*VirtualMachineWhereInput `json:"hasVirtualMachineWith,omitempty"`
}

// A connection to a list of items.
type VirtualMachineConnection struct {
	// A list of edges.
	Edges []*VirtualMachineEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type VirtualMachineEdge struct {
	// The item at the end of the edge.
	Node *VirtualMachine `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for VirtualMachine connections
type VirtualMachineOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order VirtualMachines.
	Field VirtualMachineOrderField `json:"field"`
}

// VirtualMachineWhereInput is used for filtering VirtualMachine objects.
// Input was generated by ent.
type VirtualMachineWhereInput struct {
	Not *VirtualMachineWhereInput   `json:"not,omitempty"`
	And []*VirtualMachineWhereInput `json:"and,omitempty"`
	Or  []*VirtualMachineWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// vm_cpu_config_id field predicates
	VMCPUConfigID             *gidx.PrefixedID  `json:"vmCPUConfigID,omitempty"`
	VMCPUConfigIDNeq          *gidx.PrefixedID  `json:"vmCPUConfigIDNEQ,omitempty"`
	VMCPUConfigIDIn           []gidx.PrefixedID `json:"vmCPUConfigIDIn,omitempty"`
	VMCPUConfigIDNotIn        []gidx.PrefixedID `json:"vmCPUConfigIDNotIn,omitempty"`
	VMCPUConfigIDGt           *gidx.PrefixedID  `json:"vmCPUConfigIDGT,omitempty"`
	VMCPUConfigIDGte          *gidx.PrefixedID  `json:"vmCPUConfigIDGTE,omitempty"`
	VMCPUConfigIDLt           *gidx.PrefixedID  `json:"vmCPUConfigIDLT,omitempty"`
	VMCPUConfigIDLte          *gidx.PrefixedID  `json:"vmCPUConfigIDLTE,omitempty"`
	VMCPUConfigIDContains     *gidx.PrefixedID  `json:"vmCPUConfigIDContains,omitempty"`
	VMCPUConfigIDHasPrefix    *gidx.PrefixedID  `json:"vmCPUConfigIDHasPrefix,omitempty"`
	VMCPUConfigIDHasSuffix    *gidx.PrefixedID  `json:"vmCPUConfigIDHasSuffix,omitempty"`
	VMCPUConfigIDEqualFold    *gidx.PrefixedID  `json:"vmCPUConfigIDEqualFold,omitempty"`
	VMCPUConfigIDContainsFold *gidx.PrefixedID  `json:"vmCPUConfigIDContainsFold,omitempty"`
	// virtual_machine_cpu_config edge predicates
	HasVirtualMachineCPUConfig     *bool                                `json:"hasVirtualMachineCPUConfig,omitempty"`
	HasVirtualMachineCPUConfigWith []*VirtualMachineCPUConfigWhereInput `json:"hasVirtualMachineCPUConfigWith,omitempty"`
}

type Service struct {
	Sdl *string `json:"sdl,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which VirtualMachineCPUConfig connections can be ordered.
type VirtualMachineCPUConfigOrderField string

const (
	VirtualMachineCPUConfigOrderFieldID      VirtualMachineCPUConfigOrderField = "ID"
	VirtualMachineCPUConfigOrderFieldCores   VirtualMachineCPUConfigOrderField = "cores"
	VirtualMachineCPUConfigOrderFieldSockets VirtualMachineCPUConfigOrderField = "sockets"
)

var AllVirtualMachineCPUConfigOrderField = []VirtualMachineCPUConfigOrderField{
	VirtualMachineCPUConfigOrderFieldID,
	VirtualMachineCPUConfigOrderFieldCores,
	VirtualMachineCPUConfigOrderFieldSockets,
}

func (e VirtualMachineCPUConfigOrderField) IsValid() bool {
	switch e {
	case VirtualMachineCPUConfigOrderFieldID, VirtualMachineCPUConfigOrderFieldCores, VirtualMachineCPUConfigOrderFieldSockets:
		return true
	}
	return false
}

func (e VirtualMachineCPUConfigOrderField) String() string {
	return string(e)
}

func (e *VirtualMachineCPUConfigOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VirtualMachineCPUConfigOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VirtualMachineCPUConfigOrderField", str)
	}
	return nil
}

func (e VirtualMachineCPUConfigOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which VirtualMachine connections can be ordered.
type VirtualMachineOrderField string

const (
	VirtualMachineOrderFieldID          VirtualMachineOrderField = "ID"
	VirtualMachineOrderFieldCreatedAt   VirtualMachineOrderField = "CREATED_AT"
	VirtualMachineOrderFieldUpdatedAt   VirtualMachineOrderField = "UPDATED_AT"
	VirtualMachineOrderFieldName        VirtualMachineOrderField = "NAME"
	VirtualMachineOrderFieldOwner       VirtualMachineOrderField = "OWNER"
	VirtualMachineOrderFieldVMCPUConfig VirtualMachineOrderField = "VM_CPU_CONFIG"
)

var AllVirtualMachineOrderField = []VirtualMachineOrderField{
	VirtualMachineOrderFieldID,
	VirtualMachineOrderFieldCreatedAt,
	VirtualMachineOrderFieldUpdatedAt,
	VirtualMachineOrderFieldName,
	VirtualMachineOrderFieldOwner,
	VirtualMachineOrderFieldVMCPUConfig,
}

func (e VirtualMachineOrderField) IsValid() bool {
	switch e {
	case VirtualMachineOrderFieldID, VirtualMachineOrderFieldCreatedAt, VirtualMachineOrderFieldUpdatedAt, VirtualMachineOrderFieldName, VirtualMachineOrderFieldOwner, VirtualMachineOrderFieldVMCPUConfig:
		return true
	}
	return false
}

func (e VirtualMachineOrderField) String() string {
	return string(e)
}

func (e *VirtualMachineOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VirtualMachineOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VirtualMachineOrderField", str)
	}
	return nil
}

func (e VirtualMachineOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
