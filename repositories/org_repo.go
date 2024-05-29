package repositories

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"fmt"
	"strings"
)

type OrgRepo interface {
	GetAll(filter dtos.OrgsFilterDTO) ([]entities.OrganizationEntity, error)
	GetOne(filter dtos.OrgsFilterDTO) (*entities.OrganizationEntity, error)
}

type OrgRepoPgSQL struct {
}

func NewOrgRepoPgSQL() OrgRepo {
	return &OrgRepoPgSQL{}
}

func (or *OrgRepoPgSQL) GetAll(filter dtos.OrgsFilterDTO) ([]entities.OrganizationEntity, error) {
	queryBase := "SELECT id, name FROM organizations"

	conditions, args := or.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []entities.OrganizationEntity
	for rows.Next() {
		var organization entities.OrganizationEntity
		if err := rows.Scan(&organization.Id, &organization.Name); err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	return organizations, nil
}

func (or *OrgRepoPgSQL) GetOne(filter dtos.OrgsFilterDTO) (*entities.OrganizationEntity, error) {
	queryBase := "SELECT id, name FROM organizations"

	conditions, args := or.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organization entities.OrganizationEntity
	if rows.Next() {
		if err := rows.Scan(&organization.Id, &organization.Name); err != nil {
			return nil, err
		}
	} else {
		return nil, RecNotFound
	}

	if rows.Next() {
		return nil, MultipleRecFound
	}

	return &organization, nil
}

func (or *OrgRepoPgSQL) buildQueryParams(filter dtos.OrgsFilterDTO) ([]string, []interface{}) {
	var conditions []string
	var args []interface{}

	if len(filter.Ids) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "id", placeholders(len(filter.Ids), len(args)+1)))
		for _, arg := range filter.Ids {
			args = append(args, arg)
		}
	}

	if len(filter.Names) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "name", placeholders(len(filter.Names), len(args)+1)))
		for _, arg := range filter.Names {
			args = append(args, arg)
		}
	}

	return conditions, args
}
