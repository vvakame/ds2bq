// generated by qbg -output model_query.go -private .; DO NOT EDIT

package ds2bq

import (
	"github.com/favclip/qbg/qbgutils"
	"google.golang.org/appengine/datastore"
)

// aeDatastoreAdminOperationQueryBuilder build query for aeDatastoreAdminOperation.
type aeDatastoreAdminOperationQueryBuilder struct {
	q             *datastore.Query
	plugin        qbgutils.Plugin
	Kind          *aeDatastoreAdminOperationQueryProperty
	ID            *aeDatastoreAdminOperationQueryProperty
	ActiveJobIDs  *aeDatastoreAdminOperationQueryProperty
	ActiveJobs    *aeDatastoreAdminOperationQueryProperty
	CompletedJobs *aeDatastoreAdminOperationQueryProperty
	Description   *aeDatastoreAdminOperationQueryProperty
	LastUpdated   *aeDatastoreAdminOperationQueryProperty
	ServiceJobID  *aeDatastoreAdminOperationQueryProperty
	Status        *aeDatastoreAdminOperationQueryProperty
	StatusInfo    *aeDatastoreAdminOperationQueryProperty
}

// aeDatastoreAdminOperationQueryProperty has property information for aeDatastoreAdminOperationQueryBuilder.
type aeDatastoreAdminOperationQueryProperty struct {
	bldr *aeDatastoreAdminOperationQueryBuilder
	name string
}

// newAEDatastoreAdminOperationQueryBuilder create new AEDatastoreAdminOperationQueryBuilder.
func newAEDatastoreAdminOperationQueryBuilder() *aeDatastoreAdminOperationQueryBuilder {
	return newAEDatastoreAdminOperationQueryBuilderWithKind("_AE_DatastoreAdmin_Operation")
}

// newAEDatastoreAdminOperationQueryBuilderWithKind create new AEDatastoreAdminOperationQueryBuilder with specific kind.
func newAEDatastoreAdminOperationQueryBuilderWithKind(kind string) *aeDatastoreAdminOperationQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &aeDatastoreAdminOperationQueryBuilder{q: q}
	bldr.Kind = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "Kind",
	}
	bldr.ID = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "__key__",
	}
	bldr.ActiveJobIDs = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "active_job_ids",
	}
	bldr.ActiveJobs = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "active_jobs",
	}
	bldr.CompletedJobs = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "completed_jobs",
	}
	bldr.Description = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "description",
	}
	bldr.LastUpdated = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "last_updated",
	}
	bldr.ServiceJobID = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "service_job_id",
	}
	bldr.Status = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "status",
	}
	bldr.StatusInfo = &aeDatastoreAdminOperationQueryProperty{
		bldr: bldr,
		name: "status_info",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("AEDatastoreAdminOperation")
	}
	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) Ancestor(parentKey *datastore.Key) *aeDatastoreAdminOperationQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) KeysOnly() *aeDatastoreAdminOperationQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) Start(cur datastore.Cursor) *aeDatastoreAdminOperationQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) Offset(offset int) *aeDatastoreAdminOperationQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) Limit(limit int) *aeDatastoreAdminOperationQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *aeDatastoreAdminOperationQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *aeDatastoreAdminOperationQueryProperty) Filter(op string, value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *aeDatastoreAdminOperationQueryProperty) LessThanOrEqual(value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *aeDatastoreAdminOperationQueryProperty) GreaterThanOrEqual(value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *aeDatastoreAdminOperationQueryProperty) LessThan(value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *aeDatastoreAdminOperationQueryProperty) GreaterThan(value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *aeDatastoreAdminOperationQueryProperty) Equal(value interface{}) *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *aeDatastoreAdminOperationQueryProperty) Asc() *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *aeDatastoreAdminOperationQueryProperty) Desc() *aeDatastoreAdminOperationQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}

// aeBackupInformationQueryBuilder build query for aeBackupInformation.
type aeBackupInformationQueryBuilder struct {
	q             *datastore.Query
	plugin        qbgutils.Plugin
	Kind          *aeBackupInformationQueryProperty
	ID            *aeBackupInformationQueryProperty
	CompleteTime  *aeBackupInformationQueryProperty
	CompletedJobs *aeBackupInformationQueryProperty
	Destination   *aeBackupInformationQueryProperty
	Filesystem    *aeBackupInformationQueryProperty
	GSHandle      *aeBackupInformationQueryProperty
	Kinds         *aeBackupInformationQueryProperty
	Name          *aeBackupInformationQueryProperty
	OriginalApp   *aeBackupInformationQueryProperty
	StartTime     *aeBackupInformationQueryProperty
}

// aeBackupInformationQueryProperty has property information for aeBackupInformationQueryBuilder.
type aeBackupInformationQueryProperty struct {
	bldr *aeBackupInformationQueryBuilder
	name string
}

// newAEBackupInformationQueryBuilder create new AEBackupInformationQueryBuilder.
func newAEBackupInformationQueryBuilder() *aeBackupInformationQueryBuilder {
	return newAEBackupInformationQueryBuilderWithKind("_AE_Backup_Information")
}

// newAEBackupInformationQueryBuilderWithKind create new AEBackupInformationQueryBuilder with specific kind.
func newAEBackupInformationQueryBuilderWithKind(kind string) *aeBackupInformationQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &aeBackupInformationQueryBuilder{q: q}
	bldr.Kind = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "Kind",
	}
	bldr.ID = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "__key__",
	}
	bldr.CompleteTime = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "complete_time",
	}
	bldr.CompletedJobs = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "completed_jobs",
	}
	bldr.Destination = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "destination",
	}
	bldr.Filesystem = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "filesystem",
	}
	bldr.GSHandle = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "gs_handle",
	}
	bldr.Kinds = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "kinds",
	}
	bldr.Name = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "name",
	}
	bldr.OriginalApp = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "original_app",
	}
	bldr.StartTime = &aeBackupInformationQueryProperty{
		bldr: bldr,
		name: "start_time",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("AEBackupInformation")
	}
	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *aeBackupInformationQueryBuilder) Ancestor(parentKey *datastore.Key) *aeBackupInformationQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *aeBackupInformationQueryBuilder) KeysOnly() *aeBackupInformationQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *aeBackupInformationQueryBuilder) Start(cur datastore.Cursor) *aeBackupInformationQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *aeBackupInformationQueryBuilder) Offset(offset int) *aeBackupInformationQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *aeBackupInformationQueryBuilder) Limit(limit int) *aeBackupInformationQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *aeBackupInformationQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *aeBackupInformationQueryProperty) Filter(op string, value interface{}) *aeBackupInformationQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *aeBackupInformationQueryProperty) LessThanOrEqual(value interface{}) *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *aeBackupInformationQueryProperty) GreaterThanOrEqual(value interface{}) *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *aeBackupInformationQueryProperty) LessThan(value interface{}) *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *aeBackupInformationQueryProperty) GreaterThan(value interface{}) *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *aeBackupInformationQueryProperty) Equal(value interface{}) *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *aeBackupInformationQueryProperty) Asc() *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *aeBackupInformationQueryProperty) Desc() *aeBackupInformationQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}

// aeBackupInformationKindFilesQueryBuilder build query for aeBackupInformationKindFiles.
type aeBackupInformationKindFilesQueryBuilder struct {
	q      *datastore.Query
	plugin qbgutils.Plugin
	Kind   *aeBackupInformationKindFilesQueryProperty
	ID     *aeBackupInformationKindFilesQueryProperty
	Files  *aeBackupInformationKindFilesQueryProperty
}

// aeBackupInformationKindFilesQueryProperty has property information for aeBackupInformationKindFilesQueryBuilder.
type aeBackupInformationKindFilesQueryProperty struct {
	bldr *aeBackupInformationKindFilesQueryBuilder
	name string
}

// newAEBackupInformationKindFilesQueryBuilder create new AEBackupInformationKindFilesQueryBuilder.
func newAEBackupInformationKindFilesQueryBuilder() *aeBackupInformationKindFilesQueryBuilder {
	return newAEBackupInformationKindFilesQueryBuilderWithKind("_AE_Backup_Information_Kind_Files")
}

// newAEBackupInformationKindFilesQueryBuilderWithKind create new AEBackupInformationKindFilesQueryBuilder with specific kind.
func newAEBackupInformationKindFilesQueryBuilderWithKind(kind string) *aeBackupInformationKindFilesQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &aeBackupInformationKindFilesQueryBuilder{q: q}
	bldr.Kind = &aeBackupInformationKindFilesQueryProperty{
		bldr: bldr,
		name: "Kind",
	}
	bldr.ID = &aeBackupInformationKindFilesQueryProperty{
		bldr: bldr,
		name: "__key__",
	}
	bldr.Files = &aeBackupInformationKindFilesQueryProperty{
		bldr: bldr,
		name: "files",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("AEBackupInformationKindFiles")
	}
	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) Ancestor(parentKey *datastore.Key) *aeBackupInformationKindFilesQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) KeysOnly() *aeBackupInformationKindFilesQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) Start(cur datastore.Cursor) *aeBackupInformationKindFilesQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) Offset(offset int) *aeBackupInformationKindFilesQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) Limit(limit int) *aeBackupInformationKindFilesQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *aeBackupInformationKindFilesQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *aeBackupInformationKindFilesQueryProperty) Filter(op string, value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *aeBackupInformationKindFilesQueryProperty) LessThanOrEqual(value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *aeBackupInformationKindFilesQueryProperty) GreaterThanOrEqual(value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *aeBackupInformationKindFilesQueryProperty) LessThan(value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *aeBackupInformationKindFilesQueryProperty) GreaterThan(value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *aeBackupInformationKindFilesQueryProperty) Equal(value interface{}) *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *aeBackupInformationKindFilesQueryProperty) Asc() *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *aeBackupInformationKindFilesQueryProperty) Desc() *aeBackupInformationKindFilesQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}

// aeBackupKindQueryBuilder build query for aeBackupKind.
type aeBackupKindQueryBuilder struct {
	q      *datastore.Query
	plugin qbgutils.Plugin
	Kind   *aeBackupKindQueryProperty
	ID     *aeBackupKindQueryProperty
}

// aeBackupKindQueryProperty has property information for aeBackupKindQueryBuilder.
type aeBackupKindQueryProperty struct {
	bldr *aeBackupKindQueryBuilder
	name string
}

// newAEBackupKindQueryBuilder create new AEBackupKindQueryBuilder.
func newAEBackupKindQueryBuilder() *aeBackupKindQueryBuilder {
	return newAEBackupKindQueryBuilderWithKind("Kind")
}

// newAEBackupKindQueryBuilderWithKind create new AEBackupKindQueryBuilder with specific kind.
func newAEBackupKindQueryBuilderWithKind(kind string) *aeBackupKindQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &aeBackupKindQueryBuilder{q: q}
	bldr.Kind = &aeBackupKindQueryProperty{
		bldr: bldr,
		name: "Kind",
	}
	bldr.ID = &aeBackupKindQueryProperty{
		bldr: bldr,
		name: "__key__",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("AEBackupKind")
	}
	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *aeBackupKindQueryBuilder) Ancestor(parentKey *datastore.Key) *aeBackupKindQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *aeBackupKindQueryBuilder) KeysOnly() *aeBackupKindQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *aeBackupKindQueryBuilder) Start(cur datastore.Cursor) *aeBackupKindQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *aeBackupKindQueryBuilder) Offset(offset int) *aeBackupKindQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *aeBackupKindQueryBuilder) Limit(limit int) *aeBackupKindQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *aeBackupKindQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *aeBackupKindQueryProperty) Filter(op string, value interface{}) *aeBackupKindQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *aeBackupKindQueryProperty) LessThanOrEqual(value interface{}) *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *aeBackupKindQueryProperty) GreaterThanOrEqual(value interface{}) *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *aeBackupKindQueryProperty) LessThan(value interface{}) *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *aeBackupKindQueryProperty) GreaterThan(value interface{}) *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *aeBackupKindQueryProperty) Equal(value interface{}) *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *aeBackupKindQueryProperty) Asc() *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *aeBackupKindQueryProperty) Desc() *aeBackupKindQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}

// aeBackupInformationKindTypeInfoQueryBuilder build query for aeBackupInformationKindTypeInfo.
type aeBackupInformationKindTypeInfoQueryBuilder struct {
	q              *datastore.Query
	plugin         qbgutils.Plugin
	Kind           *aeBackupInformationKindTypeInfoQueryProperty
	ID             *aeBackupInformationKindTypeInfoQueryProperty
	EntityTypeInfo *aeBackupInformationKindTypeInfoQueryProperty
	IsPartial      *aeBackupInformationKindTypeInfoQueryProperty
}

// aeBackupInformationKindTypeInfoQueryProperty has property information for aeBackupInformationKindTypeInfoQueryBuilder.
type aeBackupInformationKindTypeInfoQueryProperty struct {
	bldr *aeBackupInformationKindTypeInfoQueryBuilder
	name string
}

// newAEBackupInformationKindTypeInfoQueryBuilder create new AEBackupInformationKindTypeInfoQueryBuilder.
func newAEBackupInformationKindTypeInfoQueryBuilder() *aeBackupInformationKindTypeInfoQueryBuilder {
	return newAEBackupInformationKindTypeInfoQueryBuilderWithKind("_AE_Backup_Information_Kind_Type_Info")
}

// newAEBackupInformationKindTypeInfoQueryBuilderWithKind create new AEBackupInformationKindTypeInfoQueryBuilder with specific kind.
func newAEBackupInformationKindTypeInfoQueryBuilderWithKind(kind string) *aeBackupInformationKindTypeInfoQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &aeBackupInformationKindTypeInfoQueryBuilder{q: q}
	bldr.Kind = &aeBackupInformationKindTypeInfoQueryProperty{
		bldr: bldr,
		name: "Kind",
	}
	bldr.ID = &aeBackupInformationKindTypeInfoQueryProperty{
		bldr: bldr,
		name: "__key__",
	}
	bldr.EntityTypeInfo = &aeBackupInformationKindTypeInfoQueryProperty{
		bldr: bldr,
		name: "entity_type_info",
	}
	bldr.IsPartial = &aeBackupInformationKindTypeInfoQueryProperty{
		bldr: bldr,
		name: "is_partial",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("AEBackupInformationKindTypeInfo")
	}
	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) Ancestor(parentKey *datastore.Key) *aeBackupInformationKindTypeInfoQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) KeysOnly() *aeBackupInformationKindTypeInfoQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) Start(cur datastore.Cursor) *aeBackupInformationKindTypeInfoQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) Offset(offset int) *aeBackupInformationKindTypeInfoQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) Limit(limit int) *aeBackupInformationKindTypeInfoQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *aeBackupInformationKindTypeInfoQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) Filter(op string, value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) LessThanOrEqual(value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) GreaterThanOrEqual(value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) LessThan(value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) GreaterThan(value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *aeBackupInformationKindTypeInfoQueryProperty) Equal(value interface{}) *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *aeBackupInformationKindTypeInfoQueryProperty) Asc() *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *aeBackupInformationKindTypeInfoQueryProperty) Desc() *aeBackupInformationKindTypeInfoQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}
