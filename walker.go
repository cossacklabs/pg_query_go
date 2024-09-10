package pg_query

// Visit defines the signature of a function that
// can be used to visit all nodes of a parse tree.
type Visit func(node *Node) (kontinue bool, err error)

type Walker interface {
	WalkSubtree(visit Visit) error
}

func WalkSubtree(node isNode_Node, visit Visit) error {
	if walker, ok := node.(Walker); ok {
		return walker.WalkSubtree(visit)
	}
	return nil
}

func (n *Node_Alias) WalkSubtree(visit Visit) error {
	if n.Alias == nil {
		return nil
	}
	return Walk(visit, n.Alias.Colnames...)
}

func (n *Node_RangeVar) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_TableFunc) WalkSubtree(visit Visit) error {
	if n.TableFunc == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TableFunc.GetNsUris()...)
	nodes = append(nodes, n.TableFunc.GetNsNames()...)
	nodes = append(nodes, n.TableFunc.GetDocexpr())
	nodes = append(nodes, n.TableFunc.GetRowexpr())
	nodes = append(nodes, n.TableFunc.GetColexprs()...)
	nodes = append(nodes, n.TableFunc.GetColnames()...)
	nodes = append(nodes, n.TableFunc.GetColtypes()...)
	nodes = append(nodes, n.TableFunc.GetColcollations()...)
	nodes = append(nodes, n.TableFunc.GetColexprs()...)
	nodes = append(nodes, n.TableFunc.GetColdefexprs()...)

	return Walk(visit, nodes...)
}

func (n *Node_Var) WalkSubtree(visit Visit) error {
	if n.Var == nil {
		return nil
	}

	return Walk(visit, n.Var.GetXpr())
}

func (n *Node_Param) WalkSubtree(visit Visit) error {
	if n.Param == nil {
		return nil
	}

	return Walk(visit, n.Param.GetXpr())
}

func (n *Node_Aggref) WalkSubtree(visit Visit) error {
	if n.Aggref == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.Aggref.GetXpr())
	nodes = append(nodes, n.Aggref.GetAggargtypes()...)
	nodes = append(nodes, n.Aggref.GetAggdirectargs()...)
	nodes = append(nodes, n.Aggref.GetArgs()...)
	nodes = append(nodes, n.Aggref.GetAggorder()...)
	nodes = append(nodes, n.Aggref.GetAggdistinct()...)
	nodes = append(nodes, n.Aggref.GetAggfilter())
	return Walk(visit, nodes...)
}

func (n *Node_GroupingFunc) WalkSubtree(visit Visit) error {
	if n.GroupingFunc == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.GroupingFunc.GetXpr())
	nodes = append(nodes, n.GroupingFunc.GetRefs()...)
	nodes = append(nodes, n.GroupingFunc.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_WindowFunc) WalkSubtree(visit Visit) error {
	if n.WindowFunc == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.WindowFunc.GetXpr())
	nodes = append(nodes, n.WindowFunc.GetArgs()...)
	nodes = append(nodes, n.WindowFunc.GetAggfilter())
	return Walk(visit, nodes...)
}

func (n *Node_SubscriptingRef) WalkSubtree(visit Visit) error {
	if n.SubscriptingRef == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SubscriptingRef.GetXpr())
	nodes = append(nodes, n.SubscriptingRef.GetRefassgnexpr())
	nodes = append(nodes, n.SubscriptingRef.GetReflowerindexpr()...)
	return Walk(visit, nodes...)

}

func (n *Node_FuncExpr) WalkSubtree(visit Visit) error {
	if n.FuncExpr == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.FuncExpr.GetXpr())
	nodes = append(nodes, n.FuncExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_NamedArgExpr) WalkSubtree(visit Visit) error {
	if n.NamedArgExpr == nil {
		return nil
	}
	return Walk(visit, n.NamedArgExpr.GetXpr(), n.NamedArgExpr.GetArg())
}
func (n *Node_OpExpr) WalkSubtree(visit Visit) error {
	if n.OpExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.OpExpr.GetXpr())
	nodes = append(nodes, n.OpExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_DistinctExpr) WalkSubtree(visit Visit) error {
	if n.DistinctExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DistinctExpr.GetXpr())
	nodes = append(nodes, n.DistinctExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_NullIfExpr) WalkSubtree(visit Visit) error {
	if n.NullIfExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.NullIfExpr.GetXpr())
	nodes = append(nodes, n.NullIfExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_ScalarArrayOpExpr) WalkSubtree(visit Visit) error {
	if n.ScalarArrayOpExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ScalarArrayOpExpr.GetXpr())
	nodes = append(nodes, n.ScalarArrayOpExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_BoolExpr) WalkSubtree(visit Visit) error {
	if n.BoolExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.BoolExpr.GetXpr())
	nodes = append(nodes, n.BoolExpr.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_SubLink) WalkSubtree(visit Visit) error {
	if n.SubLink == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SubLink.GetXpr())
	nodes = append(nodes, n.SubLink.GetSubselect())
	nodes = append(nodes, n.SubLink.GetOperName()...)
	nodes = append(nodes, n.SubLink.GetTestexpr())
	return Walk(visit, nodes...)
}

func (n *Node_SubPlan) WalkSubtree(visit Visit) error {
	if n.SubPlan == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SubPlan.GetXpr())
	nodes = append(nodes, n.SubPlan.GetTestexpr())
	nodes = append(nodes, n.SubPlan.GetArgs()...)
	nodes = append(nodes, n.SubPlan.GetParamIds()...)
	nodes = append(nodes, n.SubPlan.GetParParam()...)
	nodes = append(nodes, n.SubPlan.GetSetParam()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlternativeSubPlan) WalkSubtree(visit Visit) error {
	if n.AlternativeSubPlan == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlternativeSubPlan.GetXpr())
	nodes = append(nodes, n.AlternativeSubPlan.GetSubplans()...)
	return Walk(visit, nodes...)
}

func (n *Node_FieldSelect) WalkSubtree(visit Visit) error {
	if n.FieldSelect == nil {
		return nil
	}
	return Walk(visit, n.FieldSelect.GetArg(), n.FieldSelect.GetArg())
}

func (n *Node_FieldStore) WalkSubtree(visit Visit) error {
	if n.FieldStore == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.FieldStore.GetXpr())
	nodes = append(nodes, n.FieldStore.GetArg())
	nodes = append(nodes, n.FieldStore.GetFieldnums()...)
	nodes = append(nodes, n.FieldStore.GetNewvals()...)
	return Walk(visit, nodes...)
}

func (n *Node_RelabelType) WalkSubtree(visit Visit) error {
	if n.RelabelType == nil {
		return nil
	}

	return Walk(visit, n.RelabelType.GetXpr(), n.RelabelType.GetArg())
}

func (n *Node_CoerceViaIo) WalkSubtree(visit Visit) error {
	if n.CoerceViaIo == nil {
		return nil
	}
	return Walk(visit, n.CoerceViaIo.GetXpr(), n.CoerceViaIo.GetArg())
}

func (n *Node_ArrayCoerceExpr) WalkSubtree(visit Visit) error {
	if n.ArrayCoerceExpr == nil {
		return nil
	}
	return Walk(visit, n.ArrayCoerceExpr.GetXpr(), n.ArrayCoerceExpr.GetArg(), n.ArrayCoerceExpr.GetElemexpr())
}

func (n *Node_ConvertRowtypeExpr) WalkSubtree(visit Visit) error {
	if n.ConvertRowtypeExpr == nil {
		return nil
	}
	return Walk(visit, n.ConvertRowtypeExpr.GetXpr(), n.ConvertRowtypeExpr.GetArg())
}

func (n *Node_CollateExpr) WalkSubtree(visit Visit) error {
	if n.CollateExpr == nil {
		return nil
	}
	return Walk(visit, n.CollateExpr.GetXpr(), n.CollateExpr.GetArg())
}

func (n *Node_CaseExpr) WalkSubtree(visit Visit) error {
	if n.CaseExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CaseExpr.GetXpr())
	nodes = append(nodes, n.CaseExpr.GetArg())
	nodes = append(nodes, n.CaseExpr.GetArgs()...)
	nodes = append(nodes, n.CaseExpr.GetDefresult())
	return Walk(visit, nodes...)
}

func (n *Node_CaseWhen) WalkSubtree(visit Visit) error {
	if n.CaseWhen == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CaseWhen.GetXpr())
	nodes = append(nodes, n.CaseWhen.GetExpr())
	return Walk(visit, nodes...)
}

func (n *Node_CaseTestExpr) WalkSubtree(visit Visit) error {
	if n.CaseTestExpr == nil {
		return nil
	}
	return Walk(visit, n.CaseTestExpr.GetXpr())
}

func (n *Node_ArrayExpr) WalkSubtree(visit Visit) error {
	if n.ArrayExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ArrayExpr.GetXpr())
	nodes = append(nodes, n.ArrayExpr.GetElements()...)
	return Walk(visit, nodes...)
}

func (n *Node_RowExpr) WalkSubtree(visit Visit) error {
	if n.RowExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RowExpr.GetXpr())
	nodes = append(nodes, n.RowExpr.GetArgs()...)
	nodes = append(nodes, n.RowExpr.GetColnames()...)
	return Walk(visit, nodes...)
}

func (n *Node_RowCompareExpr) WalkSubtree(visit Visit) error {
	if n.RowCompareExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RowCompareExpr.GetXpr())
	nodes = append(nodes, n.RowCompareExpr.GetInputcollids()...)
	nodes = append(nodes, n.RowCompareExpr.GetLargs()...)
	nodes = append(nodes, n.RowCompareExpr.GetOpfamilies()...)
	nodes = append(nodes, n.RowCompareExpr.GetOpnos()...)
	nodes = append(nodes, n.RowCompareExpr.GetRargs()...)
	return Walk(visit, nodes...)
}

func (n *Node_CoalesceExpr) WalkSubtree(visit Visit) error {
	if n.CoalesceExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CoalesceExpr.GetArgs()...)
	nodes = append(nodes, n.CoalesceExpr.GetXpr())
	return Walk(visit, nodes...)
}

func (n *Node_MinMaxExpr) WalkSubtree(visit Visit) error {
	if n.MinMaxExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.MinMaxExpr.GetArgs()...)
	nodes = append(nodes, n.MinMaxExpr.GetXpr())
	return Walk(visit, nodes...)
}

func (n *Node_SqlvalueFunction) WalkSubtree(visit Visit) error {
	if n.SqlvalueFunction == nil {
		return nil
	}
	return Walk(visit, n.SqlvalueFunction.GetXpr())
}

func (n *Node_XmlExpr) WalkSubtree(visit Visit) error {
	if n.XmlExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.XmlExpr.GetArgs()...)
	nodes = append(nodes, n.XmlExpr.GetXpr())
	nodes = append(nodes, n.XmlExpr.GetArgNames()...)
	nodes = append(nodes, n.XmlExpr.GetNamedArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_NullTest) WalkSubtree(visit Visit) error {
	if n.NullTest == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.NullTest.GetArg())
	nodes = append(nodes, n.NullTest.GetXpr())
	return Walk(visit, nodes...)
}

func (n *Node_BooleanTest) WalkSubtree(visit Visit) error {
	if n.BooleanTest == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.BooleanTest.GetArg())
	nodes = append(nodes, n.BooleanTest.GetXpr())
	return Walk(visit, nodes...)
}

func (n *Node_CoerceToDomain) WalkSubtree(visit Visit) error {
	if n.CoerceToDomain == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CoerceToDomain.GetArg())
	nodes = append(nodes, n.CoerceToDomain.GetXpr())
	return Walk(visit, nodes...)
}

func (n *Node_CoerceToDomainValue) WalkSubtree(visit Visit) error {
	if n.CoerceToDomainValue == nil {
		return nil
	}
	return Walk(visit, n.CoerceToDomainValue.GetXpr())
}

func (n *Node_SetToDefault) WalkSubtree(visit Visit) error {
	if n.SetToDefault == nil {
		return nil
	}
	return Walk(visit, n.SetToDefault.GetXpr())
}

func (n *Node_CurrentOfExpr) WalkSubtree(visit Visit) error {
	if n.CurrentOfExpr == nil {
		return nil
	}
	return Walk(visit, n.CurrentOfExpr.GetXpr())
}

func (n *Node_NextValueExpr) WalkSubtree(visit Visit) error {
	if n.NextValueExpr == nil {
		return nil
	}
	return Walk(visit, n.NextValueExpr.GetXpr())
}

func (n *Node_InferenceElem) WalkSubtree(visit Visit) error {
	if n.InferenceElem == nil {
		return nil
	}
	return Walk(visit, n.InferenceElem.GetExpr(), n.InferenceElem.GetXpr())
}

func (n *Node_TargetEntry) WalkSubtree(visit Visit) error {
	if n.TargetEntry == nil {
		return nil
	}
	return Walk(visit, n.TargetEntry.GetExpr(), n.TargetEntry.GetXpr())
}

func (n *Node_RangeTblRef) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_JoinExpr) WalkSubtree(visit Visit) error {
	if n.JoinExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.JoinExpr.GetRarg())
	nodes = append(nodes, n.JoinExpr.GetLarg())
	nodes = append(nodes, n.JoinExpr.GetUsingClause()...)
	nodes = append(nodes, n.JoinExpr.GetQuals())
	return Walk(visit, nodes...)
}

func (n *Node_FromExpr) WalkSubtree(visit Visit) error {
	if n.FromExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.FromExpr.GetQuals())
	nodes = append(nodes, n.FromExpr.GetFromlist()...)
	return Walk(visit, nodes...)
}

func (n *Node_OnConflictExpr) WalkSubtree(visit Visit) error {
	if n.OnConflictExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.OnConflictExpr.GetArbiterElems()...)
	nodes = append(nodes, n.OnConflictExpr.GetExclRelTlist()...)
	nodes = append(nodes, n.OnConflictExpr.GetArbiterElems()...)
	nodes = append(nodes, n.OnConflictExpr.GetOnConflictWhere())
	return Walk(visit, nodes...)
}

func (n *Node_IntoClause) WalkSubtree(visit Visit) error {
	if n.IntoClause == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.IntoClause.GetColNames()...)
	nodes = append(nodes, n.IntoClause.GetOptions()...)
	nodes = append(nodes, n.IntoClause.GetViewQuery())
	return Walk(visit, nodes...)
}

func (n *Node_MergeAction) WalkSubtree(visit Visit) error {
	if n.MergeAction == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.MergeAction.GetTargetList()...)
	nodes = append(nodes, n.MergeAction.GetUpdateColnos()...)
	nodes = append(nodes, n.MergeAction.GetQual())
	return Walk(visit, nodes...)
}

func (n *Node_RawStmt) WalkSubtree(visit Visit) error {
	if n.RawStmt == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RawStmt.GetStmt())
	return Walk(visit, nodes...)
}

func (n *Node_Query) WalkSubtree(visit Visit) error {
	if n.Query == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.Query.GetTargetList()...)
	nodes = append(nodes, n.Query.GetConstraintDeps()...)
	nodes = append(nodes, n.Query.GetCteList()...)
	nodes = append(nodes, n.Query.GetRtable()...)
	nodes = append(nodes, n.Query.GetGroupClause()...)
	nodes = append(nodes, n.Query.GetGroupingSets()...)
	nodes = append(nodes, n.Query.GetSortClause()...)
	nodes = append(nodes, n.Query.GetMergeActionList()...)
	nodes = append(nodes, n.Query.GetDistinctClause()...)
	nodes = append(nodes, n.Query.GetLimitCount())
	nodes = append(nodes, n.Query.GetHavingQual())
	nodes = append(nodes, n.Query.GetLimitOffset())
	nodes = append(nodes, n.Query.GetRowMarks()...)
	nodes = append(nodes, n.Query.GetReturningList()...)
	nodes = append(nodes, n.Query.GetWithCheckOptions()...)
	nodes = append(nodes, n.Query.GetWindowClause()...)
	nodes = append(nodes, n.Query.GetSetOperations())
	nodes = append(nodes, n.Query.GetUtilityStmt())
	return Walk(visit, nodes...)
}

func (n *Node_InsertStmt) WalkSubtree(visit Visit) error {
	if n.InsertStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.InsertStmt.GetCols()...)
	nodes = append(nodes, n.InsertStmt.GetSelectStmt())
	nodes = append(nodes, n.InsertStmt.GetReturningList()...)
	return Walk(visit, nodes...)
}

func (n *Node_DeleteStmt) WalkSubtree(visit Visit) error {
	if n.DeleteStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DeleteStmt.GetUsingClause()...)
	nodes = append(nodes, n.DeleteStmt.GetWhereClause())
	nodes = append(nodes, n.DeleteStmt.GetReturningList()...)
	return Walk(visit, nodes...)
}

func (n *Node_UpdateStmt) WalkSubtree(visit Visit) error {
	if n.UpdateStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.UpdateStmt.GetTargetList()...)
	nodes = append(nodes, n.UpdateStmt.GetFromClause()...)
	nodes = append(nodes, n.UpdateStmt.GetWhereClause())
	nodes = append(nodes, n.UpdateStmt.GetReturningList()...)
	return Walk(visit, nodes...)
}

func (n *Node_MergeStmt) WalkSubtree(visit Visit) error {
	if n.MergeStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.MergeStmt.GetJoinCondition())
	nodes = append(nodes, n.MergeStmt.GetMergeWhenClauses()...)
	nodes = append(nodes, n.MergeStmt.GetSourceRelation())
	return Walk(visit, nodes...)
}

func (n *Node_SelectStmt) WalkSubtree(visit Visit) error {
	if n.SelectStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SelectStmt.GetTargetList()...)
	nodes = append(nodes, n.SelectStmt.GetFromClause()...)
	nodes = append(nodes, n.SelectStmt.GetWhereClause())
	nodes = append(nodes, n.SelectStmt.GetSortClause()...)
	nodes = append(nodes, n.SelectStmt.GetLimitOffset())
	nodes = append(nodes, n.SelectStmt.GetWindowClause()...)
	nodes = append(nodes, n.SelectStmt.GetLockingClause()...)
	nodes = append(nodes, n.SelectStmt.GetValuesLists()...)
	nodes = append(nodes, n.SelectStmt.GetGroupClause()...)
	nodes = append(nodes, n.SelectStmt.GetDistinctClause()...)
	nodes = append(nodes, n.SelectStmt.GetHavingClause())
	nodes = append(nodes, n.SelectStmt.GetLimitCount())
	return Walk(visit, nodes...)
}

func (n *Node_ReturnStmt) WalkSubtree(visit Visit) error {
	if n.ReturnStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ReturnStmt.GetReturnval())
	return Walk(visit, nodes...)
}

func (n *Node_PlassignStmt) WalkSubtree(visit Visit) error {
	if n.PlassignStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.PlassignStmt.GetIndirection()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTableStmt) WalkSubtree(visit Visit) error {
	if n.AlterTableStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTableStmt.GetCmds()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTableCmd) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterDomainStmt) WalkSubtree(visit Visit) error {
	if n.AlterDomainStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterDomainStmt.GetTypeName()...)
	return Walk(visit, nodes...)
}

func (n *Node_SetOperationStmt) WalkSubtree(visit Visit) error {
	if n.SetOperationStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SetOperationStmt.GetRarg())
	nodes = append(nodes, n.SetOperationStmt.GetLarg())
	nodes = append(nodes, n.SetOperationStmt.GetColCollations()...)
	nodes = append(nodes, n.SetOperationStmt.GetColTypes()...)
	nodes = append(nodes, n.SetOperationStmt.GetColTypmods()...)
	nodes = append(nodes, n.SetOperationStmt.GetGroupClauses()...)
	return Walk(visit, nodes...)
}

func (n *Node_GrantStmt) WalkSubtree(visit Visit) error {
	if n.GrantStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.GrantStmt.GetGrantees()...)
	nodes = append(nodes, n.GrantStmt.GetObjects()...)
	nodes = append(nodes, n.GrantStmt.GetPrivileges()...)
	return Walk(visit, nodes...)
}

func (n *Node_GrantRoleStmt) WalkSubtree(visit Visit) error {
	if n.GrantRoleStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.GrantRoleStmt.GetGrantedRoles()...)
	nodes = append(nodes, n.GrantRoleStmt.GetGranteeRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterDefaultPrivilegesStmt) WalkSubtree(visit Visit) error {
	if n.AlterDefaultPrivilegesStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterDefaultPrivilegesStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_ClosePortalStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_ClusterStmt) WalkSubtree(visit Visit) error {
	if n.ClusterStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ClusterStmt.GetParams()...)
	return Walk(visit, nodes...)
}

func (n *Node_CopyStmt) WalkSubtree(visit Visit) error {
	if n.CopyStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CopyStmt.GetOptions()...)
	nodes = append(nodes, n.CopyStmt.GetWhereClause())
	nodes = append(nodes, n.CopyStmt.GetAttlist()...)
	nodes = append(nodes, n.CopyStmt.GetQuery())
	return Walk(visit, nodes...)
}

func (n *Node_CreateStmt) WalkSubtree(visit Visit) error {
	if n.CreateStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateStmt.GetOptions()...)
	nodes = append(nodes, n.CreateStmt.GetConstraints()...)
	nodes = append(nodes, n.CreateStmt.GetInhRelations()...)
	nodes = append(nodes, n.CreateStmt.GetTableElts()...)
	return Walk(visit, nodes...)
}

func (n *Node_DefineStmt) WalkSubtree(visit Visit) error {
	if n.DefineStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DefineStmt.GetArgs()...)
	nodes = append(nodes, n.DefineStmt.GetDefnames()...)
	nodes = append(nodes, n.DefineStmt.GetDefinition()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropStmt) WalkSubtree(visit Visit) error {
	if n.DropStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DropStmt.GetObjects()...)
	return Walk(visit, nodes...)
}

func (n *Node_TruncateStmt) WalkSubtree(visit Visit) error {
	if n.TruncateStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TruncateStmt.GetRelations()...)
	return Walk(visit, nodes...)
}

func (n *Node_CommentStmt) WalkSubtree(visit Visit) error {
	if n.CommentStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CommentStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_FetchStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_IndexStmt) WalkSubtree(visit Visit) error {
	if n.IndexStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.IndexStmt.GetOptions()...)
	nodes = append(nodes, n.IndexStmt.GetWhereClause())
	nodes = append(nodes, n.IndexStmt.GetExcludeOpNames()...)
	nodes = append(nodes, n.IndexStmt.GetIndexParams()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateFunctionStmt) WalkSubtree(visit Visit) error {
	if n.CreateFunctionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateFunctionStmt.GetOptions()...)
	nodes = append(nodes, n.CreateFunctionStmt.GetFuncname()...)
	nodes = append(nodes, n.CreateFunctionStmt.GetParameters()...)
	nodes = append(nodes, n.CreateFunctionStmt.GetSqlBody())
	return Walk(visit, nodes...)
}

func (n *Node_AlterFunctionStmt) WalkSubtree(visit Visit) error {
	if n.AlterFunctionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterFunctionStmt.GetActions()...)
	return Walk(visit, nodes...)
}

func (n *Node_DoStmt) WalkSubtree(visit Visit) error {
	if n.DoStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DoStmt.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_RenameStmt) WalkSubtree(visit Visit) error {
	if n.RenameStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RenameStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_RuleStmt) WalkSubtree(visit Visit) error {
	if n.RuleStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RuleStmt.GetActions()...)
	nodes = append(nodes, n.RuleStmt.GetWhereClause())
	return Walk(visit, nodes...)
}

func (n *Node_NotifyStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_ListenStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_UnlistenStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_TransactionStmt) WalkSubtree(visit Visit) error {
	if n.TransactionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TransactionStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_ViewStmt) WalkSubtree(visit Visit) error {
	if n.ViewStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ViewStmt.GetQuery())
	nodes = append(nodes, n.ViewStmt.GetOptions()...)
	nodes = append(nodes, n.ViewStmt.GetAliases()...)
	return Walk(visit, nodes...)
}

func (n *Node_LoadStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateDomainStmt) WalkSubtree(visit Visit) error {
	if n.CreateDomainStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateDomainStmt.GetConstraints()...)
	nodes = append(nodes, n.CreateDomainStmt.GetDomainname()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreatedbStmt) WalkSubtree(visit Visit) error {
	if n.CreatedbStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreatedbStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropdbStmt) WalkSubtree(visit Visit) error {
	if n.DropdbStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DropdbStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_VacuumStmt) WalkSubtree(visit Visit) error {
	if n.VacuumStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.VacuumStmt.GetOptions()...)
	nodes = append(nodes, n.VacuumStmt.GetRels()...)
	return Walk(visit, nodes...)
}

func (n *Node_ExplainStmt) WalkSubtree(visit Visit) error {
	if n.ExplainStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ExplainStmt.GetQuery())
	nodes = append(nodes, n.ExplainStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateTableAsStmt) WalkSubtree(visit Visit) error {
	if n.CreateTableAsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateTableAsStmt.GetQuery())
	return Walk(visit, nodes...)
}

func (n *Node_CreateSeqStmt) WalkSubtree(visit Visit) error {
	if n.CreateSeqStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateSeqStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterSeqStmt) WalkSubtree(visit Visit) error {
	if n.AlterSeqStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterSeqStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_VariableSetStmt) WalkSubtree(visit Visit) error {
	if n.VariableSetStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.VariableSetStmt.GetArgs()...)
	return Walk(visit, nodes...)
}

func (n *Node_VariableShowStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_DiscardStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateTrigStmt) WalkSubtree(visit Visit) error {
	if n.CreateTrigStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateTrigStmt.GetArgs()...)
	nodes = append(nodes, n.CreateTrigStmt.GetFuncname()...)
	nodes = append(nodes, n.CreateTrigStmt.GetFuncname()...)
	nodes = append(nodes, n.CreateTrigStmt.GetTransitionRels()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreatePlangStmt) WalkSubtree(visit Visit) error {
	if n.CreatePlangStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreatePlangStmt.GetPlinline()...)
	nodes = append(nodes, n.CreatePlangStmt.GetPlhandler()...)
	nodes = append(nodes, n.CreatePlangStmt.GetPlvalidator()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateRoleStmt) WalkSubtree(visit Visit) error {
	if n.CreateRoleStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateRoleStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterRoleStmt) WalkSubtree(visit Visit) error {
	if n.AlterRoleStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterRoleStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropRoleStmt) WalkSubtree(visit Visit) error {
	if n.DropRoleStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DropRoleStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_LockStmt) WalkSubtree(visit Visit) error {
	if n.LockStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.LockStmt.GetRelations()...)
	return Walk(visit, nodes...)
}

func (n *Node_ConstraintsSetStmt) WalkSubtree(visit Visit) error {
	if n.ConstraintsSetStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ConstraintsSetStmt.GetConstraints()...)
	return Walk(visit, nodes...)
}

func (n *Node_ReindexStmt) WalkSubtree(visit Visit) error {
	if n.ReindexStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ReindexStmt.GetParams()...)
	return Walk(visit, nodes...)
}

func (n *Node_CheckPointStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateSchemaStmt) WalkSubtree(visit Visit) error {
	if n.CreateSchemaStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateSchemaStmt.GetSchemaElts()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterDatabaseStmt) WalkSubtree(visit Visit) error {
	if n.AlterDatabaseStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterDatabaseStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterDatabaseRefreshCollStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterDatabaseSetStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterRoleSetStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateConversionStmt) WalkSubtree(visit Visit) error {
	if n.CreateConversionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateConversionStmt.GetConversionName()...)
	nodes = append(nodes, n.CreateConversionStmt.GetFuncName()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateCastStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateOpClassStmt) WalkSubtree(visit Visit) error {
	if n.CreateOpClassStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateOpClassStmt.GetItems()...)
	nodes = append(nodes, n.CreateOpClassStmt.GetOpclassname()...)
	nodes = append(nodes, n.CreateOpClassStmt.GetOpfamilyname()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateOpFamilyStmt) WalkSubtree(visit Visit) error {
	if n.CreateOpFamilyStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateOpFamilyStmt.GetOpfamilyname()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterOpFamilyStmt) WalkSubtree(visit Visit) error {
	if n.AlterOpFamilyStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterOpFamilyStmt.GetItems()...)
	nodes = append(nodes, n.AlterOpFamilyStmt.GetOpfamilyname()...)
	return Walk(visit, nodes...)
}

func (n *Node_PrepareStmt) WalkSubtree(visit Visit) error {
	if n.PrepareStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.PrepareStmt.GetArgtypes()...)
	nodes = append(nodes, n.PrepareStmt.GetQuery())
	return Walk(visit, nodes...)
}

func (n *Node_ExecuteStmt) WalkSubtree(visit Visit) error {
	if n.ExecuteStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ExecuteStmt.GetParams()...)
	return Walk(visit, nodes...)
}

func (n *Node_DeallocateStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_DeclareCursorStmt) WalkSubtree(visit Visit) error {
	if n.DeclareCursorStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DeclareCursorStmt.GetQuery())
	return Walk(visit, nodes...)
}

func (n *Node_CreateTableSpaceStmt) WalkSubtree(visit Visit) error {
	if n.CreateTableSpaceStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateTableSpaceStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropTableSpaceStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterObjectDependsStmt) WalkSubtree(visit Visit) error {
	if n.AlterObjectDependsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterObjectDependsStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_AlterObjectSchemaStmt) WalkSubtree(visit Visit) error {
	if n.AlterObjectSchemaStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterObjectSchemaStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_AlterOwnerStmt) WalkSubtree(visit Visit) error {
	if n.AlterOwnerStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterOwnerStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_AlterOperatorStmt) WalkSubtree(visit Visit) error {
	if n.AlterOperatorStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterOperatorStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTypeStmt) WalkSubtree(visit Visit) error {
	if n.AlterTypeStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTypeStmt.GetOptions()...)
	nodes = append(nodes, n.AlterTypeStmt.GetTypeName()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropOwnedStmt) WalkSubtree(visit Visit) error {
	if n.DropOwnedStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DropOwnedStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_ReassignOwnedStmt) WalkSubtree(visit Visit) error {
	if n.ReassignOwnedStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ReassignOwnedStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_CompositeTypeStmt) WalkSubtree(visit Visit) error {
	if n.CompositeTypeStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CompositeTypeStmt.GetColdeflist()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateEnumStmt) WalkSubtree(visit Visit) error {
	if n.CreateEnumStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateEnumStmt.GetTypeName()...)
	nodes = append(nodes, n.CreateEnumStmt.GetTypeName()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateRangeStmt) WalkSubtree(visit Visit) error {
	if n.CreateRangeStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateRangeStmt.GetTypeName()...)
	nodes = append(nodes, n.CreateRangeStmt.GetTypeName()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterEnumStmt) WalkSubtree(visit Visit) error {
	if n.AlterEnumStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterEnumStmt.GetTypeName()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTsdictionaryStmt) WalkSubtree(visit Visit) error {
	if n.AlterTsdictionaryStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTsdictionaryStmt.GetOptions()...)
	nodes = append(nodes, n.AlterTsdictionaryStmt.GetDictname()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTsconfigurationStmt) WalkSubtree(visit Visit) error {
	if n.AlterTsconfigurationStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTsconfigurationStmt.GetDicts()...)
	nodes = append(nodes, n.AlterTsconfigurationStmt.GetCfgname()...)
	nodes = append(nodes, n.AlterTsconfigurationStmt.GetTokentype()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateFdwStmt) WalkSubtree(visit Visit) error {
	if n.CreateFdwStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateFdwStmt.GetFuncOptions()...)
	nodes = append(nodes, n.CreateFdwStmt.GetFuncOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterFdwStmt) WalkSubtree(visit Visit) error {
	if n.AlterFdwStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterFdwStmt.GetOptions()...)
	nodes = append(nodes, n.AlterFdwStmt.GetFuncOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateForeignServerStmt) WalkSubtree(visit Visit) error {
	if n.CreateForeignServerStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateForeignServerStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterForeignServerStmt) WalkSubtree(visit Visit) error {
	if n.AlterForeignServerStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterForeignServerStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateUserMappingStmt) WalkSubtree(visit Visit) error {
	if n.CreateUserMappingStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateUserMappingStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterUserMappingStmt) WalkSubtree(visit Visit) error {
	if n.AlterUserMappingStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterUserMappingStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropUserMappingStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterTableSpaceOptionsStmt) WalkSubtree(visit Visit) error {
	if n.AlterTableSpaceOptionsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTableSpaceOptionsStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterTableMoveAllStmt) WalkSubtree(visit Visit) error {
	if n.AlterTableMoveAllStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterTableMoveAllStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_SecLabelStmt) WalkSubtree(visit Visit) error {
	if n.SecLabelStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SecLabelStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_CreateForeignTableStmt) WalkSubtree(visit Visit) error {
	if n.CreateForeignTableStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateForeignTableStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_ImportForeignSchemaStmt) WalkSubtree(visit Visit) error {
	if n.ImportForeignSchemaStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ImportForeignSchemaStmt.GetOptions()...)
	nodes = append(nodes, n.ImportForeignSchemaStmt.GetTableList()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateExtensionStmt) WalkSubtree(visit Visit) error {
	if n.CreateExtensionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateExtensionStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterExtensionStmt) WalkSubtree(visit Visit) error {
	if n.AlterExtensionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterExtensionStmt.GetOptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterExtensionContentsStmt) WalkSubtree(visit Visit) error {
	if n.AlterExtensionContentsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterExtensionContentsStmt.GetObject())
	return Walk(visit, nodes...)
}

func (n *Node_CreateEventTrigStmt) WalkSubtree(visit Visit) error {
	if n.CreateEventTrigStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateEventTrigStmt.GetFuncname()...)
	nodes = append(nodes, n.CreateEventTrigStmt.GetWhenclause()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterEventTrigStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_RefreshMatViewStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_ReplicaIdentityStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AlterSystemStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreatePolicyStmt) WalkSubtree(visit Visit) error {
	if n.CreatePolicyStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreatePolicyStmt.GetQual())
	nodes = append(nodes, n.CreatePolicyStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterPolicyStmt) WalkSubtree(visit Visit) error {
	if n.AlterPolicyStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterPolicyStmt.GetQual())
	nodes = append(nodes, n.AlterPolicyStmt.GetRoles()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateTransformStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateAmStmt) WalkSubtree(visit Visit) error {
	if n.CreateAmStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateAmStmt.GetHandlerName()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreatePublicationStmt) WalkSubtree(visit Visit) error {
	if n.CreatePublicationStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreatePublicationStmt.GetOptions()...)
	nodes = append(nodes, n.CreatePublicationStmt.GetPubobjects()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterPublicationStmt) WalkSubtree(visit Visit) error {
	if n.AlterPublicationStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterPublicationStmt.GetOptions()...)
	nodes = append(nodes, n.AlterPublicationStmt.GetPubobjects()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateSubscriptionStmt) WalkSubtree(visit Visit) error {
	if n.CreateSubscriptionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateSubscriptionStmt.GetOptions()...)
	nodes = append(nodes, n.CreateSubscriptionStmt.GetPublication()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterSubscriptionStmt) WalkSubtree(visit Visit) error {
	if n.AlterSubscriptionStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterSubscriptionStmt.GetOptions()...)
	nodes = append(nodes, n.AlterSubscriptionStmt.GetPublication()...)
	return Walk(visit, nodes...)
}

func (n *Node_DropSubscriptionStmt) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CreateStatsStmt) WalkSubtree(visit Visit) error {
	if n.CreateStatsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateStatsStmt.GetRelations()...)
	nodes = append(nodes, n.CreateStatsStmt.GetDefnames()...)
	nodes = append(nodes, n.CreateStatsStmt.GetExprs()...)
	nodes = append(nodes, n.CreateStatsStmt.GetStatTypes()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterCollationStmt) WalkSubtree(visit Visit) error {
	if n.AlterCollationStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterCollationStmt.GetCollname()...)
	return Walk(visit, nodes...)
}

func (n *Node_CallStmt) WalkSubtree(visit Visit) error {
	if n.CallStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CallStmt.GetOutargs()...)
	return Walk(visit, nodes...)
}

func (n *Node_AlterStatsStmt) WalkSubtree(visit Visit) error {
	if n.AlterStatsStmt == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AlterStatsStmt.GetDefnames()...)
	return Walk(visit, nodes...)
}

func (n *Node_AExpr) WalkSubtree(visit Visit) error {
	if n.AExpr == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AExpr.GetName()...)
	return Walk(visit, nodes...)
}

func (n *Node_ColumnRef) WalkSubtree(visit Visit) error {
	if n.ColumnRef == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ColumnRef.GetFields()...)
	return Walk(visit, nodes...)
}

func (n *Node_ParamRef) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_FuncCall) WalkSubtree(visit Visit) error {
	if n.FuncCall == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.FuncCall.GetArgs()...)
	nodes = append(nodes, n.FuncCall.GetFuncname()...)
	nodes = append(nodes, n.FuncCall.GetAggOrder()...)
	nodes = append(nodes, n.FuncCall.GetAggFilter())
	return Walk(visit, nodes...)
}

func (n *Node_AStar) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_AIndices) WalkSubtree(visit Visit) error {
	if n.AIndices == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AIndices.GetLidx())
	nodes = append(nodes, n.AIndices.GetUidx())
	return Walk(visit, nodes...)
}

func (n *Node_AIndirection) WalkSubtree(visit Visit) error {
	if n.AIndirection == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AIndirection.GetIndirection()...)
	nodes = append(nodes, n.AIndirection.GetArg())
	return Walk(visit, nodes...)
}

func (n *Node_AArrayExpr) WalkSubtree(visit Visit) error {
	if n.AArrayExpr == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AArrayExpr.GetElements()...)
	return Walk(visit, nodes...)
}

func (n *Node_ResTarget) WalkSubtree(visit Visit) error {
	if n.ResTarget == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ResTarget.GetIndirection()...)
	nodes = append(nodes, n.ResTarget.GetVal())
	return Walk(visit, nodes...)
}

func (n *Node_MultiAssignRef) WalkSubtree(visit Visit) error {
	if n.MultiAssignRef == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.MultiAssignRef.GetSource())
	return Walk(visit, nodes...)
}

func (n *Node_TypeCast) WalkSubtree(visit Visit) error {
	if n.TypeCast == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TypeCast.GetArg())
	return Walk(visit, nodes...)
}

func (n *Node_CollateClause) WalkSubtree(visit Visit) error {
	if n.CollateClause == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CollateClause.GetArg())
	nodes = append(nodes, n.CollateClause.GetCollname()...)
	return Walk(visit, nodes...)
}

func (n *Node_SortBy) WalkSubtree(visit Visit) error {
	if n.SortBy == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.SortBy.GetNode())
	nodes = append(nodes, n.SortBy.GetUseOp()...)
	return Walk(visit, nodes...)
}

func (n *Node_WindowDef) WalkSubtree(visit Visit) error {
	if n.WindowDef == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.WindowDef.GetPartitionClause()...)
	nodes = append(nodes, n.WindowDef.GetPartitionClause()...)
	nodes = append(nodes, n.WindowDef.GetStartOffset())
	nodes = append(nodes, n.WindowDef.GetStartOffset())
	return Walk(visit, nodes...)
}

func (n *Node_RangeSubselect) WalkSubtree(visit Visit) error {
	if n.RangeSubselect == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeSubselect.GetSubquery())
	return Walk(visit, nodes...)
}

func (n *Node_RangeFunction) WalkSubtree(visit Visit) error {
	if n.RangeFunction == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeFunction.GetFunctions()...)
	nodes = append(nodes, n.RangeFunction.GetColdeflist()...)
	return Walk(visit, nodes...)
}

func (n *Node_RangeTableSample) WalkSubtree(visit Visit) error {
	if n.RangeTableSample == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeTableSample.GetArgs()...)
	nodes = append(nodes, n.RangeTableSample.GetRepeatable())
	nodes = append(nodes, n.RangeTableSample.GetRelation())
	nodes = append(nodes, n.RangeTableSample.GetMethod()...)
	return Walk(visit, nodes...)
}

func (n *Node_RangeTableFunc) WalkSubtree(visit Visit) error {
	if n.RangeTableFunc == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeTableFunc.GetDocexpr())
	nodes = append(nodes, n.RangeTableFunc.GetRowexpr())
	nodes = append(nodes, n.RangeTableFunc.GetColumns()...)
	nodes = append(nodes, n.RangeTableFunc.GetNamespaces()...)
	return Walk(visit, nodes...)
}

func (n *Node_RangeTableFuncCol) WalkSubtree(visit Visit) error {
	if n.RangeTableFuncCol == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeTableFuncCol.GetColdefexpr())
	nodes = append(nodes, n.RangeTableFuncCol.GetColexpr())
	return Walk(visit, nodes...)
}

func (n *Node_TypeName) WalkSubtree(visit Visit) error {
	if n.TypeName == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TypeName.GetNames()...)
	nodes = append(nodes, n.TypeName.GetArrayBounds()...)
	nodes = append(nodes, n.TypeName.GetTypmods()...)
	return Walk(visit, nodes...)
}

func (n *Node_ColumnDef) WalkSubtree(visit Visit) error {
	if n.ColumnDef == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ColumnDef.GetConstraints()...)
	nodes = append(nodes, n.ColumnDef.GetCookedDefault())
	nodes = append(nodes, n.ColumnDef.GetFdwoptions()...)
	return Walk(visit, nodes...)
}

func (n *Node_IndexElem) WalkSubtree(visit Visit) error {
	if n.IndexElem == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.IndexElem.GetExpr())
	nodes = append(nodes, n.IndexElem.GetOpclass()...)
	nodes = append(nodes, n.IndexElem.GetCollation()...)
	nodes = append(nodes, n.IndexElem.GetOpclassopts()...)
	return Walk(visit, nodes...)
}

func (n *Node_StatsElem) WalkSubtree(visit Visit) error {
	if n.StatsElem == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.StatsElem.GetExpr())
	return Walk(visit, nodes...)
}

func (n *Node_Constraint) WalkSubtree(visit Visit) error {
	if n.Constraint == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.Constraint.GetOptions()...)
	nodes = append(nodes, n.Constraint.GetExclusions()...)
	nodes = append(nodes, n.Constraint.GetFkAttrs()...)
	nodes = append(nodes, n.Constraint.GetFkDelSetCols()...)
	nodes = append(nodes, n.Constraint.GetRawExpr())
	nodes = append(nodes, n.Constraint.GetIncluding()...)
	return Walk(visit, nodes...)
}

func (n *Node_DefElem) WalkSubtree(visit Visit) error {
	if n.DefElem == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.DefElem.GetArg())
	return Walk(visit, nodes...)
}

func (n *Node_RangeTblEntry) WalkSubtree(visit Visit) error {
	if n.RangeTblEntry == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeTblEntry.GetColtypes()...)
	nodes = append(nodes, n.RangeTblEntry.GetColcollations()...)
	nodes = append(nodes, n.RangeTblEntry.GetColtypes()...)
	nodes = append(nodes, n.RangeTblEntry.GetColtypmods()...)
	nodes = append(nodes, n.RangeTblEntry.GetFunctions()...)
	nodes = append(nodes, n.RangeTblEntry.GetSecurityQuals()...)
	nodes = append(nodes, n.RangeTblEntry.GetJoinrightcols()...)
	nodes = append(nodes, n.RangeTblEntry.GetJoinaliasvars()...)
	return Walk(visit, nodes...)
}

func (n *Node_RangeTblFunction) WalkSubtree(visit Visit) error {
	if n.RangeTblFunction == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.RangeTblFunction.GetFunccolnames()...)
	nodes = append(nodes, n.RangeTblFunction.GetFunccolcollations()...)
	nodes = append(nodes, n.RangeTblFunction.GetFunccoltypes()...)
	nodes = append(nodes, n.RangeTblFunction.GetFunccoltypmods()...)
	nodes = append(nodes, n.RangeTblFunction.GetFuncexpr())
	return Walk(visit, nodes...)
}

func (n *Node_TableSampleClause) WalkSubtree(visit Visit) error {
	if n.TableSampleClause == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.TableSampleClause.GetArgs()...)
	nodes = append(nodes, n.TableSampleClause.GetRepeatable())
	return Walk(visit, nodes...)
}

func (n *Node_WithCheckOption) WalkSubtree(visit Visit) error {
	if n.WithCheckOption == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.WithCheckOption.GetQual())
	return Walk(visit, nodes...)
}

func (n *Node_SortGroupClause) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_GroupingSet) WalkSubtree(visit Visit) error {
	if n.GroupingSet == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.GroupingSet.GetContent()...)
	return Walk(visit, nodes...)
}

func (n *Node_WindowClause) WalkSubtree(visit Visit) error {
	if n.WindowClause == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.WindowClause.GetEndOffset())
	nodes = append(nodes, n.WindowClause.GetOrderClause()...)
	nodes = append(nodes, n.WindowClause.GetPartitionClause()...)
	nodes = append(nodes, n.WindowClause.GetRunCondition()...)
	nodes = append(nodes, n.WindowClause.GetStartOffset())
	return Walk(visit, nodes...)
}

func (n *Node_ObjectWithArgs) WalkSubtree(visit Visit) error {
	if n.ObjectWithArgs == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.ObjectWithArgs.GetObjargs()...)
	nodes = append(nodes, n.ObjectWithArgs.GetObjfuncargs()...)
	nodes = append(nodes, n.ObjectWithArgs.GetObjname()...)
	return Walk(visit, nodes...)
}

func (n *Node_AccessPriv) WalkSubtree(visit Visit) error {
	if n.AccessPriv == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.AccessPriv.GetCols()...)
	return Walk(visit, nodes...)
}

func (n *Node_CreateOpClassItem) WalkSubtree(visit Visit) error {
	if n.CreateOpClassItem == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CreateOpClassItem.GetClassArgs()...)
	nodes = append(nodes, n.CreateOpClassItem.GetOrderFamily()...)
	return Walk(visit, nodes...)
}

func (n *Node_TableLikeClause) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_FunctionParameter) WalkSubtree(visit Visit) error {
	if n.FunctionParameter == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.FunctionParameter.GetDefexpr())
	return Walk(visit, nodes...)
}

func (n *Node_LockingClause) WalkSubtree(visit Visit) error {
	if n.LockingClause == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.LockingClause.GetLockedRels()...)
	return Walk(visit, nodes...)
}

func (n *Node_RowMarkClause) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_XmlSerialize) WalkSubtree(visit Visit) error {
	if n.XmlSerialize == nil {
		return nil
	}

	nodes := make([]*Node, 0)
	nodes = append(nodes, n.XmlSerialize.GetExpr())
	return Walk(visit, nodes...)
}

func (n *Node_WithClause) WalkSubtree(visit Visit) error {
	if n.WithClause == nil {
		return nil
	}
	return Walk(visit, n.WithClause.GetCtes()...)
}

func (n *Node_InferClause) WalkSubtree(visit Visit) error {
	if n.InferClause == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.InferClause.GetWhereClause())
	nodes = append(nodes, n.InferClause.GetIndexElems()...)
	return Walk(visit, nodes...)
}

func (n *Node_OnConflictClause) WalkSubtree(visit Visit) error {
	if n.OnConflictClause == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.OnConflictClause.GetTargetList()...)
	nodes = append(nodes, n.OnConflictClause.GetWhereClause())
	return Walk(visit, nodes...)
}

func (n *Node_CtesearchClause) WalkSubtree(visit Visit) error {
	if n.CtesearchClause == nil {
		return nil
	}
	return Walk(visit, n.CtesearchClause.GetSearchColList()...)
}

func (n *Node_CtecycleClause) WalkSubtree(visit Visit) error {
	if n.CtecycleClause == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CtecycleClause.GetCycleColList()...)
	nodes = append(nodes, n.CtecycleClause.GetCycleMarkDefault())
	return Walk(visit, nodes...)

}

func (n *Node_CommonTableExpr) WalkSubtree(visit Visit) error {
	if n.CommonTableExpr == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.CommonTableExpr.GetCtecolcollations()...)
	nodes = append(nodes, n.CommonTableExpr.GetCtecolnames()...)
	nodes = append(nodes, n.CommonTableExpr.GetAliascolnames()...)
	nodes = append(nodes, n.CommonTableExpr.GetCtecoltypes()...)
	nodes = append(nodes, n.CommonTableExpr.GetCtecoltypmods()...)
	nodes = append(nodes, n.CommonTableExpr.GetCtequery())
	return Walk(visit, nodes...)
}

func (n *Node_MergeWhenClause) WalkSubtree(visit Visit) error {
	if n.MergeWhenClause == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.MergeWhenClause.GetTargetList()...)
	nodes = append(nodes, n.MergeWhenClause.GetValues()...)
	return Walk(visit, nodes...)
}

func (n *Node_RoleSpec) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_TriggerTransition) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_PartitionElem) WalkSubtree(visit Visit) error {
	if n.PartitionElem == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.PartitionElem.GetCollation()...)
	nodes = append(nodes, n.PartitionElem.GetOpclass()...)
	nodes = append(nodes, n.PartitionElem.GetExpr())
	return Walk(visit, nodes...)
}

func (n *Node_PartitionSpec) WalkSubtree(visit Visit) error {
	if n.PartitionSpec == nil {
		return nil
	}
	return Walk(visit, n.PartitionSpec.GetPartParams()...)
}

func (n *Node_PartitionBoundSpec) WalkSubtree(visit Visit) error {
	if n.PartitionBoundSpec == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.PartitionBoundSpec.GetListdatums()...)
	nodes = append(nodes, n.PartitionBoundSpec.GetLowerdatums()...)
	nodes = append(nodes, n.PartitionBoundSpec.GetUpperdatums()...)
	return Walk(visit, nodes...)
}

func (n *Node_PartitionRangeDatum) WalkSubtree(visit Visit) error {
	if n.PartitionRangeDatum == nil {
		return nil
	}
	return Walk(visit, n.PartitionRangeDatum.GetValue())
}

func (n *Node_PartitionCmd) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_VacuumRelation) WalkSubtree(visit Visit) error {
	if n.VacuumRelation == nil {
		return nil
	}
	return Walk(visit, n.VacuumRelation.GetVaCols()...)
}

func (n *Node_PublicationObjSpec) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_PublicationTable) WalkSubtree(visit Visit) error {
	if n.PublicationTable == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes = append(nodes, n.PublicationTable.GetWhereClause())
	nodes = append(nodes, n.PublicationTable.GetColumns()...)
	return Walk(visit, nodes...)
}

func (n *Node_InlineCodeBlock) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_CallContext) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_Integer) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_Float) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_Boolean) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_String_) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_BitString) WalkSubtree(visit Visit) error {
	return nil
}

func (n *Node_List) WalkSubtree(visit Visit) error {
	if n.List == nil {
		return nil
	}
	return Walk(visit, n.List.GetItems()...)
}

func (n *Node_IntList) WalkSubtree(visit Visit) error {
	if n.IntList == nil {
		return nil
	}
	return Walk(visit, n.IntList.GetItems()...)
}

func (n *Node_OidList) WalkSubtree(visit Visit) error {
	if n.OidList == nil {
		return nil
	}
	return Walk(visit, n.OidList.GetItems()...)
}

func (n *Node_AConst) WalkSubtree(visit Visit) error {
	return nil
}
