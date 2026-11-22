// Code generated from ../EXPRESS/t.exp. DO NOT EDIT.

package ifc

var entityExpressDocs = map[string]string{
	"IFCACTIONREQUEST": `ENTITY IfcActionRequest
 SUBTYPE OF (IfcControl);
	PredefinedType : OPTIONAL IfcActionRequestTypeEnum;
	Status : OPTIONAL IfcLabel;
	LongDescription : OPTIONAL IfcText;
END_ENTITY;`,
	"IFCACTOR": `ENTITY IfcActor
 SUPERTYPE OF (ONEOF
	(IfcOccupant))
 SUBTYPE OF (IfcObject);
	TheActor : IfcActorSelect;
 INVERSE
	IsActingUpon : SET [0:?] OF IfcRelAssignsToActor FOR RelatingActor;
END_ENTITY;`,
	"IFCACTORROLE": `ENTITY IfcActorRole;
	Role : IfcRoleEnum;
	UserDefinedRole : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
 INVERSE
	HasExternalReference : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
 WHERE
	WR1 : (Role <> IfcRoleEnum.USERDEFINED) OR
((Role = IfcRoleEnum.USERDEFINED) AND
  EXISTS(SELF.UserDefinedRole));
END_ENTITY;`,
	"IFCACTUATOR": `ENTITY IfcActuator
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcActuatorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcActuatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcActuatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCACTUATORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCACTUATORTYPE": `ENTITY IfcActuatorType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcActuatorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcActuatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcActuatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCADDRESS": `ENTITY IfcAddress
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPostalAddress
	,IfcTelecomAddress));
	Purpose : OPTIONAL IfcAddressTypeEnum;
	Description : OPTIONAL IfcText;
	UserDefinedPurpose : OPTIONAL IfcLabel;
 INVERSE
	OfPerson : SET [0:?] OF IfcPerson FOR Addresses;
	OfOrganization : SET [0:?] OF IfcOrganization FOR Addresses;
 WHERE
	WR1 : (NOT(EXISTS(Purpose))) OR
((Purpose <> IfcAddressTypeEnum.USERDEFINED) OR
((Purpose = IfcAddressTypeEnum.USERDEFINED) AND
  EXISTS(SELF.UserDefinedPurpose)));
END_ENTITY;`,
	"IFCADVANCEDBREP": `ENTITY IfcAdvancedBrep
 SUPERTYPE OF (ONEOF
	(IfcAdvancedBrepWithVoids))
 SUBTYPE OF (IfcManifoldSolidBrep);
 WHERE
	HasAdvancedFaces : SIZEOF(QUERY(Afs <* SELF\IfcManifoldSolidBrep.Outer.CfsFaces |
  (NOT ('IFC4X3_DEV_73740fe4.IFCADVANCEDFACE' IN TYPEOF(Afs)))
)) = 0;
END_ENTITY;`,
	"IFCADVANCEDBREPWITHVOIDS": `ENTITY IfcAdvancedBrepWithVoids
 SUBTYPE OF (IfcAdvancedBrep);
	Voids : SET [1:?] OF IfcClosedShell;
 WHERE
	VoidsHaveAdvancedFaces : SIZEOF (QUERY (Vsh <* Voids |
  SIZEOF (QUERY (Afs <* Vsh.CfsFaces |
  (NOT ('IFC4X3_DEV_73740fe4.IFCADVANCEDFACE' IN TYPEOF(Afs)))
  )) = 0
)) = 0;
END_ENTITY;`,
	"IFCADVANCEDFACE": `ENTITY IfcAdvancedFace
 SUBTYPE OF (IfcFaceSurface);
 WHERE
	ApplicableEdgeCurves : SIZEOF(QUERY (ElpFbnds <* 
  QUERY (Bnds <* SELF\IfcFace.Bounds |
    'IFC4X3_DEV_73740fe4.IFCEDGELOOP' IN TYPEOF(Bnds.Bound)) |
     NOT (SIZEOF (QUERY (Oe <* ElpFbnds.Bound\IfcEdgeLoop.EdgeList |
       NOT (SIZEOF (['IFC4X3_DEV_73740fe4.IFCLINE',
                     'IFC4X3_DEV_73740fe4.IFCCONIC',
                     'IFC4X3_DEV_73740fe4.IFCPOLYLINE',
                     'IFC4X3_DEV_73740fe4.IFCBSPLINECURVE'] *
         TYPEOF(Oe\IfcOrientedEdge.EdgeElement\IfcEdgeCurve.EdgeGeometry)) = 1 )
  )) = 0
))) = 0;
	ApplicableSurface : SIZEOF (
['IFC4X3_DEV_73740fe4.IFCELEMENTARYSURFACE',
 'IFC4X3_DEV_73740fe4.IFCSWEPTSURFACE',
 'IFC4X3_DEV_73740fe4.IFCBSPLINESURFACE'] *
TYPEOF(SELF\IfcFaceSurface.FaceSurface)) = 1;
	RequiresEdgeCurve : SIZEOF(QUERY (ElpFbnds <* 
  QUERY (Bnds <* SELF\IfcFace.Bounds |
    'IFC4X3_DEV_73740fe4.IFCEDGELOOP' IN TYPEOF(Bnds.Bound)) |
     NOT (SIZEOF (QUERY (Oe <* ElpFbnds.Bound\IfcEdgeLoop.EdgeList |
       NOT('IFC4X3_DEV_73740fe4.IFCEDGECURVE' IN 
       TYPEOF(Oe\IfcOrientedEdge.EdgeElement)
  ))) = 0
))) = 0;
END_ENTITY;`,
	"IFCAIRTERMINAL": `ENTITY IfcAirTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcAirTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcAirTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCAIRTERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCAIRTERMINALBOX": `ENTITY IfcAirTerminalBox
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcAirTerminalBoxTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcAirTerminalBoxTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirTerminalBoxTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCAIRTERMINALBOXTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCAIRTERMINALBOXTYPE": `ENTITY IfcAirTerminalBoxType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcAirTerminalBoxTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAirTerminalBoxTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirTerminalBoxTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCAIRTERMINALTYPE": `ENTITY IfcAirTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcAirTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAirTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCAIRTOAIRHEATRECOVERY": `ENTITY IfcAirToAirHeatRecovery
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcAirToAirHeatRecoveryTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcAirToAirHeatRecoveryTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirToAirHeatRecoveryTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCAIRTOAIRHEATRECOVERYTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCAIRTOAIRHEATRECOVERYTYPE": `ENTITY IfcAirToAirHeatRecoveryType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcAirToAirHeatRecoveryTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAirToAirHeatRecoveryTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAirToAirHeatRecoveryTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCALARM": `ENTITY IfcAlarm
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcAlarmTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcAlarmTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAlarmTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCALARMTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCALARMTYPE": `ENTITY IfcAlarmType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcAlarmTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAlarmTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAlarmTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCALIGNMENT": `ENTITY IfcAlignment
 SUBTYPE OF (IfcLinearPositioningElement);
	PredefinedType : OPTIONAL IfcAlignmentTypeEnum;
END_ENTITY;`,
	"IFCALIGNMENTCANT": `ENTITY IfcAlignmentCant
 SUBTYPE OF (IfcLinearElement);
	RailHeadDistance : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCALIGNMENTCANTSEGMENT": `ENTITY IfcAlignmentCantSegment
 SUBTYPE OF (IfcAlignmentParameterSegment);
	StartDistAlong : IfcLengthMeasure;
	HorizontalLength : IfcNonNegativeLengthMeasure;
	StartCantLeft : IfcLengthMeasure;
	EndCantLeft : OPTIONAL IfcLengthMeasure;
	StartCantRight : IfcLengthMeasure;
	EndCantRight : OPTIONAL IfcLengthMeasure;
	PredefinedType : IfcAlignmentCantSegmentTypeEnum;
END_ENTITY;`,
	"IFCALIGNMENTHORIZONTAL": `ENTITY IfcAlignmentHorizontal
 SUBTYPE OF (IfcLinearElement);
END_ENTITY;`,
	"IFCALIGNMENTHORIZONTALSEGMENT": `ENTITY IfcAlignmentHorizontalSegment
 SUBTYPE OF (IfcAlignmentParameterSegment);
	StartPoint : IfcCartesianPoint;
	StartDirection : IfcPlaneAngleMeasure;
	StartRadiusOfCurvature : IfcLengthMeasure;
	EndRadiusOfCurvature : IfcLengthMeasure;
	SegmentLength : IfcNonNegativeLengthMeasure;
	GravityCenterLineHeight : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : IfcAlignmentHorizontalSegmentTypeEnum;
END_ENTITY;`,
	"IFCALIGNMENTPARAMETERSEGMENT": `ENTITY IfcAlignmentParameterSegment
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAlignmentCantSegment
	,IfcAlignmentHorizontalSegment
	,IfcAlignmentVerticalSegment));
	StartTag : OPTIONAL IfcLabel;
	EndTag : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCALIGNMENTSEGMENT": `ENTITY IfcAlignmentSegment
 SUBTYPE OF (IfcLinearElement);
	DesignParameters : IfcAlignmentParameterSegment;
END_ENTITY;`,
	"IFCALIGNMENTVERTICAL": `ENTITY IfcAlignmentVertical
 SUBTYPE OF (IfcLinearElement);
END_ENTITY;`,
	"IFCALIGNMENTVERTICALSEGMENT": `ENTITY IfcAlignmentVerticalSegment
 SUBTYPE OF (IfcAlignmentParameterSegment);
	StartDistAlong : IfcLengthMeasure;
	HorizontalLength : IfcNonNegativeLengthMeasure;
	StartHeight : IfcLengthMeasure;
	StartGradient : IfcRatioMeasure;
	EndGradient : IfcRatioMeasure;
	RadiusOfCurvature : OPTIONAL IfcLengthMeasure;
	PredefinedType : IfcAlignmentVerticalSegmentTypeEnum;
END_ENTITY;`,
	"IFCANNOTATION": `ENTITY IfcAnnotation
 SUBTYPE OF (IfcProduct);
	PredefinedType : OPTIONAL IfcAnnotationTypeEnum;
 INVERSE
	ContainedInStructure : SET [0:1] OF IfcRelContainedInSpatialStructure FOR RelatedElements;
END_ENTITY;`,
	"IFCANNOTATIONFILLAREA": `ENTITY IfcAnnotationFillArea
 SUBTYPE OF (IfcGeometricRepresentationItem);
	OuterBoundary : IfcCurve;
	InnerBoundaries : OPTIONAL SET [1:?] OF IfcCurve;
END_ENTITY;`,
	"IFCAPPLICATION": `ENTITY IfcApplication;
	ApplicationDeveloper : IfcOrganization;
	Version : IfcLabel;
	ApplicationFullName : IfcLabel;
	ApplicationIdentifier : IfcIdentifier;
 UNIQUE
	UR1 : ApplicationIdentifier;
	UR2 : ApplicationFullName, Version;
END_ENTITY;`,
	"IFCAPPLIEDVALUE": `ENTITY IfcAppliedValue
 SUPERTYPE OF (ONEOF
	(IfcCostValue));
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	AppliedValue : OPTIONAL IfcAppliedValueSelect;
	UnitBasis : OPTIONAL IfcMeasureWithUnit;
	ApplicableDate : OPTIONAL IfcDate;
	FixedUntilDate : OPTIONAL IfcDate;
	Category : OPTIONAL IfcLabel;
	Condition : OPTIONAL IfcLabel;
	ArithmeticOperator : OPTIONAL IfcArithmeticOperatorEnum;
	Components : OPTIONAL LIST [1:?] OF IfcAppliedValue;
 INVERSE
	HasExternalReference : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCAPPROVAL": `ENTITY IfcApproval;
	Identifier : OPTIONAL IfcIdentifier;
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	TimeOfApproval : OPTIONAL IfcDateTime;
	Status : OPTIONAL IfcLabel;
	Level : OPTIONAL IfcLabel;
	Qualifier : OPTIONAL IfcText;
	RequestingApproval : OPTIONAL IfcActorSelect;
	GivingApproval : OPTIONAL IfcActorSelect;
 INVERSE
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
	ApprovedObjects : SET [0:?] OF IfcRelAssociatesApproval FOR RelatingApproval;
	ApprovedResources : SET [0:?] OF IfcResourceApprovalRelationship FOR RelatingApproval;
	IsRelatedWith : SET [0:?] OF IfcApprovalRelationship FOR RelatedApprovals;
	Relates : SET [0:?] OF IfcApprovalRelationship FOR RelatingApproval;
 WHERE
	HasIdentifierOrName : EXISTS (Identifier) OR EXISTS (Name);
END_ENTITY;`,
	"IFCAPPROVALRELATIONSHIP": `ENTITY IfcApprovalRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingApproval : IfcApproval;
	RelatedApprovals : SET [1:?] OF IfcApproval;
END_ENTITY;`,
	"IFCARBITRARYCLOSEDPROFILEDEF": `ENTITY IfcArbitraryClosedProfileDef
 SUPERTYPE OF (ONEOF
	(IfcArbitraryProfileDefWithVoids))
 SUBTYPE OF (IfcProfileDef);
	OuterCurve : IfcCurve;
 WHERE
	WR1 : OuterCurve.Dim = 2;
	WR2 : NOT('IFC4X3_DEV_73740fe4.IFCLINE' IN TYPEOF(OuterCurve));
	WR3 : NOT('IFC4X3_DEV_73740fe4.IFCOFFSETCURVE2D' IN TYPEOF(OuterCurve));
END_ENTITY;`,
	"IFCARBITRARYOPENPROFILEDEF": `ENTITY IfcArbitraryOpenProfileDef
 SUPERTYPE OF (ONEOF
	(IfcCenterLineProfileDef))
 SUBTYPE OF (IfcProfileDef);
	Curve : IfcBoundedCurve;
 WHERE
	WR11 : ('IFC4X3_DEV_73740fe4.IFCCENTERLINEPROFILEDEF' IN TYPEOF(SELF)) OR 
 (SELF\IfcProfileDef.ProfileType = IfcProfileTypeEnum.CURVE);
	WR12 : Curve.Dim = 2;
END_ENTITY;`,
	"IFCARBITRARYPROFILEDEFWITHVOIDS": `ENTITY IfcArbitraryProfileDefWithVoids
 SUBTYPE OF (IfcArbitraryClosedProfileDef);
	InnerCurves : SET [1:?] OF IfcCurve;
 WHERE
	WR1 : SELF\IfcProfileDef.ProfileType = AREA;
	WR2 : SIZEOF(QUERY(temp <* InnerCurves | temp.Dim <> 2)) = 0;
	WR3 : SIZEOF(QUERY(temp <* InnerCurves | 'IFC4X3_DEV_73740fe4.IFCLINE' IN TYPEOF(temp))) = 0;
END_ENTITY;`,
	"IFCASSET": `ENTITY IfcAsset
 SUBTYPE OF (IfcGroup);
	Identification : OPTIONAL IfcIdentifier;
	OriginalValue : OPTIONAL IfcCostValue;
	CurrentValue : OPTIONAL IfcCostValue;
	TotalReplacementCost : OPTIONAL IfcCostValue;
	Owner : OPTIONAL IfcActorSelect;
	User : OPTIONAL IfcActorSelect;
	ResponsiblePerson : OPTIONAL IfcPerson;
	IncorporationDate : OPTIONAL IfcDate;
	DepreciatedValue : OPTIONAL IfcCostValue;
END_ENTITY;`,
	"IFCASYMMETRICISHAPEPROFILEDEF": `ENTITY IfcAsymmetricIShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	BottomFlangeWidth : IfcPositiveLengthMeasure;
	OverallDepth : IfcPositiveLengthMeasure;
	WebThickness : IfcPositiveLengthMeasure;
	BottomFlangeThickness : IfcPositiveLengthMeasure;
	BottomFlangeFilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	TopFlangeWidth : IfcPositiveLengthMeasure;
	TopFlangeThickness : OPTIONAL IfcPositiveLengthMeasure;
	TopFlangeFilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	BottomFlangeEdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	BottomFlangeSlope : OPTIONAL IfcPlaneAngleMeasure;
	TopFlangeEdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	TopFlangeSlope : OPTIONAL IfcPlaneAngleMeasure;
 WHERE
	ValidBottomFilletRadius : (NOT(EXISTS(BottomFlangeFilletRadius))) OR 
(BottomFlangeFilletRadius <= (BottomFlangeWidth - WebThickness)/2.);
	ValidFlangeThickness : NOT(EXISTS(TopFlangeThickness)) OR ((BottomFlangeThickness + TopFlangeThickness) < OverallDepth);
	ValidTopFilletRadius : (NOT(EXISTS(TopFlangeFilletRadius))) OR 
(TopFlangeFilletRadius <= (TopFlangeWidth - WebThickness)/2.);
	ValidWebThickness : (WebThickness < BottomFlangeWidth) AND (WebThickness < TopFlangeWidth);
END_ENTITY;`,
	"IFCAUDIOVISUALAPPLIANCE": `ENTITY IfcAudioVisualAppliance
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcAudioVisualApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcAudioVisualApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAudioVisualApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCAUDIOVISUALAPPLIANCETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCAUDIOVISUALAPPLIANCETYPE": `ENTITY IfcAudioVisualApplianceType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcAudioVisualApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAudioVisualApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAudioVisualApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCAXIS1PLACEMENT": `ENTITY IfcAxis1Placement
 SUBTYPE OF (IfcPlacement);
	Axis : OPTIONAL IfcDirection;
 DERIVE
	 Z : IfcDirection := NVL (IfcNormalise(Axis), IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcDirection([0.0,0.0,1.0]));
 WHERE
	AxisIs3D : (NOT (EXISTS (Axis))) OR (Axis.Dim = 3);
	LocationIs3D : SELF\IfcPlacement.Location.Dim = 3;
	LocationIsCP : 'IFC4X3_DEV_73740fe4.IFCCARTESIANPOINT' IN TYPEOF(SELF\IfcPlacement.Location);
END_ENTITY;`,
	"IFCAXIS2PLACEMENT2D": `ENTITY IfcAxis2Placement2D
 SUBTYPE OF (IfcPlacement);
	RefDirection : OPTIONAL IfcDirection;
 DERIVE
	 P : LIST [2:2] OF IfcDirection := IfcBuild2Axes(RefDirection);
 WHERE
	LocationIs2D : SELF\IfcPlacement.Location.Dim = 2;
	LocationIsCP : 'IFC4X3_DEV_73740fe4.IFCCARTESIANPOINT' IN TYPEOF(SELF\IfcPlacement.Location);
	RefDirIs2D : (NOT (EXISTS (RefDirection))) OR (RefDirection.Dim = 2);
END_ENTITY;`,
	"IFCAXIS2PLACEMENT3D": `ENTITY IfcAxis2Placement3D
 SUBTYPE OF (IfcPlacement);
	Axis : OPTIONAL IfcDirection;
	RefDirection : OPTIONAL IfcDirection;
 DERIVE
	 P : LIST [3:3] OF IfcDirection := IfcBuildAxes(Axis, RefDirection);
 WHERE
	AxisAndRefDirProvision : NOT ((EXISTS (Axis)) XOR (EXISTS (RefDirection)));
	AxisIs3D : (NOT (EXISTS (Axis))) OR (Axis.Dim = 3);
	AxisToRefDirPosition : (NOT (EXISTS (Axis))) OR (NOT (EXISTS (RefDirection))) OR (IfcCrossProduct(Axis,RefDirection).Magnitude > 0.0);
	LocationIs3D : SELF\IfcPlacement.Location.Dim = 3;
	LocationIsCP : 'IFC4X3_DEV_73740fe4.IFCCARTESIANPOINT' IN TYPEOF(SELF\IfcPlacement.Location);
	RefDirIs3D : (NOT (EXISTS (RefDirection))) OR (RefDirection.Dim = 3);
END_ENTITY;`,
	"IFCAXIS2PLACEMENTLINEAR": `ENTITY IfcAxis2PlacementLinear
 SUBTYPE OF (IfcPlacement);
	Axis : OPTIONAL IfcDirection;
	RefDirection : OPTIONAL IfcDirection;
 WHERE
	WR1 : 'IFC4X3_DEV_73740fe4.IFCPOINTBYDISTANCEEXPRESSION' IN TYPEOF(SELF\IfcPlacement.Location);
	WR2 : (NOT (EXISTS (Axis))) OR (NOT (EXISTS (RefDirection))) OR (IfcCrossProduct(Axis,RefDirection).Magnitude > 0.0);
END_ENTITY;`,
	"IFCBEAM": `ENTITY IfcBeam
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcBeamTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBeamTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBeamTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBEAMTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCBEAMTYPE": `ENTITY IfcBeamType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcBeamTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBeamTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBeamTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCBEARING": `ENTITY IfcBearing
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcBearingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBearingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBearingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBEARINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCBEARINGTYPE": `ENTITY IfcBearingType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcBearingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBearingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBearingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCBLOBTEXTURE": `ENTITY IfcBlobTexture
 SUBTYPE OF (IfcSurfaceTexture);
	RasterFormat : IfcIdentifier;
	RasterCode : IfcBinary;
 WHERE
	RasterCodeByteStream : BLENGTH(RasterCode) MOD 8 = 0;
	SupportedRasterFormat : SELF.RasterFormat IN ['BMP', 'JPG', 'GIF', 'PNG'];
END_ENTITY;`,
	"IFCBLOCK": `ENTITY IfcBlock
 SUBTYPE OF (IfcCsgPrimitive3D);
	XLength : IfcPositiveLengthMeasure;
	YLength : IfcPositiveLengthMeasure;
	ZLength : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCBOILER": `ENTITY IfcBoiler
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcBoilerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBoilerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBoilerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBOILERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCBOILERTYPE": `ENTITY IfcBoilerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcBoilerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBoilerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBoilerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCBOOLEANCLIPPINGRESULT": `ENTITY IfcBooleanClippingResult
 SUBTYPE OF (IfcBooleanResult);
 WHERE
	FirstOperandType : ('IFC4X3_DEV_73740fe4.IFCSWEPTAREASOLID' IN TYPEOF(FirstOperand)) OR 
('IFC4X3_DEV_73740fe4.IFCSWEPTDISCSOLID' IN TYPEOF(FirstOperand)) OR 
('IFC4X3_DEV_73740fe4.IFCBOOLEANCLIPPINGRESULT' IN TYPEOF(FirstOperand));
	OperatorType : Operator = DIFFERENCE;
	SecondOperandType : ('IFC4X3_DEV_73740fe4.IFCHALFSPACESOLID' IN TYPEOF(SecondOperand));
END_ENTITY;`,
	"IFCBOOLEANRESULT": `ENTITY IfcBooleanResult
 SUPERTYPE OF (ONEOF
	(IfcBooleanClippingResult))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Operator : IfcBooleanOperator;
	FirstOperand : IfcBooleanOperand;
	SecondOperand : IfcBooleanOperand;
 DERIVE
	 Dim : IfcDimensionCount := FirstOperand.Dim;
 WHERE
	FirstOperandClosed : NOT('IFC4X3_DEV_73740fe4.IFCTESSELLATEDFACESET' IN TYPEOF(FirstOperand)) OR (EXISTS(FirstOperand.Closed) AND FirstOperand.Closed);
	SameDim : FirstOperand.Dim = SecondOperand.Dim;
	SecondOperandClosed : NOT('IFC4X3_DEV_73740fe4.IFCTESSELLATEDFACESET' IN TYPEOF(SecondOperand)) OR (EXISTS(SecondOperand.Closed) AND SecondOperand.Closed);
END_ENTITY;`,
	"IFCBOREHOLE": `ENTITY IfcBorehole
 SUBTYPE OF (IfcGeotechnicalAssembly);
END_ENTITY;`,
	"IFCBOUNDARYCONDITION": `ENTITY IfcBoundaryCondition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBoundaryEdgeCondition
	,IfcBoundaryFaceCondition
	,IfcBoundaryNodeCondition));
	Name : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCBOUNDARYCURVE": `ENTITY IfcBoundaryCurve
 SUPERTYPE OF (ONEOF
	(IfcOuterBoundaryCurve))
 SUBTYPE OF (IfcCompositeCurveOnSurface);
 WHERE
	IsClosed : SELF\IfcCompositeCurve.ClosedCurve;
END_ENTITY;`,
	"IFCBOUNDARYEDGECONDITION": `ENTITY IfcBoundaryEdgeCondition
 SUBTYPE OF (IfcBoundaryCondition);
	TranslationalStiffnessByLengthX : OPTIONAL IfcModulusOfTranslationalSubgradeReactionSelect;
	TranslationalStiffnessByLengthY : OPTIONAL IfcModulusOfTranslationalSubgradeReactionSelect;
	TranslationalStiffnessByLengthZ : OPTIONAL IfcModulusOfTranslationalSubgradeReactionSelect;
	RotationalStiffnessByLengthX : OPTIONAL IfcModulusOfRotationalSubgradeReactionSelect;
	RotationalStiffnessByLengthY : OPTIONAL IfcModulusOfRotationalSubgradeReactionSelect;
	RotationalStiffnessByLengthZ : OPTIONAL IfcModulusOfRotationalSubgradeReactionSelect;
END_ENTITY;`,
	"IFCBOUNDARYFACECONDITION": `ENTITY IfcBoundaryFaceCondition
 SUBTYPE OF (IfcBoundaryCondition);
	TranslationalStiffnessByAreaX : OPTIONAL IfcModulusOfSubgradeReactionSelect;
	TranslationalStiffnessByAreaY : OPTIONAL IfcModulusOfSubgradeReactionSelect;
	TranslationalStiffnessByAreaZ : OPTIONAL IfcModulusOfSubgradeReactionSelect;
END_ENTITY;`,
	"IFCBOUNDARYNODECONDITION": `ENTITY IfcBoundaryNodeCondition
 SUPERTYPE OF (ONEOF
	(IfcBoundaryNodeConditionWarping))
 SUBTYPE OF (IfcBoundaryCondition);
	TranslationalStiffnessX : OPTIONAL IfcTranslationalStiffnessSelect;
	TranslationalStiffnessY : OPTIONAL IfcTranslationalStiffnessSelect;
	TranslationalStiffnessZ : OPTIONAL IfcTranslationalStiffnessSelect;
	RotationalStiffnessX : OPTIONAL IfcRotationalStiffnessSelect;
	RotationalStiffnessY : OPTIONAL IfcRotationalStiffnessSelect;
	RotationalStiffnessZ : OPTIONAL IfcRotationalStiffnessSelect;
END_ENTITY;`,
	"IFCBOUNDARYNODECONDITIONWARPING": `ENTITY IfcBoundaryNodeConditionWarping
 SUBTYPE OF (IfcBoundaryNodeCondition);
	WarpingStiffness : OPTIONAL IfcWarpingStiffnessSelect;
END_ENTITY;`,
	"IFCBOUNDEDCURVE": `ENTITY IfcBoundedCurve
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBSplineCurve
	,IfcCompositeCurve
	,IfcIndexedPolyCurve
	,IfcPolyline
	,IfcTrimmedCurve))
 SUBTYPE OF (IfcCurve);
END_ENTITY;`,
	"IFCBOUNDEDSURFACE": `ENTITY IfcBoundedSurface
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBSplineSurface
	,IfcCurveBoundedPlane
	,IfcCurveBoundedSurface
	,IfcRectangularTrimmedSurface))
 SUBTYPE OF (IfcSurface);
END_ENTITY;`,
	"IFCBOUNDINGBOX": `ENTITY IfcBoundingBox
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Corner : IfcCartesianPoint;
	XDim : IfcPositiveLengthMeasure;
	YDim : IfcPositiveLengthMeasure;
	ZDim : IfcPositiveLengthMeasure;
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCBOXEDHALFSPACE": `ENTITY IfcBoxedHalfSpace
 SUBTYPE OF (IfcHalfSpaceSolid);
	Enclosure : IfcBoundingBox;
 WHERE
	UnboundedSurface : NOT ('IFC4X3_DEV_73740fe4.IFCCURVEBOUNDEDPLANE' IN TYPEOF(SELF\IfcHalfSpaceSolid.BaseSurface));
END_ENTITY;`,
	"IFCBRIDGE": `ENTITY IfcBridge
 SUBTYPE OF (IfcFacility);
	PredefinedType : OPTIONAL IfcBridgeTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBridgeTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBridgeTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCBRIDGEPART": `ENTITY IfcBridgePart
 SUBTYPE OF (IfcFacilityPart);
	PredefinedType : OPTIONAL IfcBridgePartTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBridgePartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBridgePartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCBSPLINECURVE": `ENTITY IfcBSplineCurve
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBSplineCurveWithKnots))
 SUBTYPE OF (IfcBoundedCurve);
	Degree : IfcInteger;
	ControlPointsList : LIST [2:?] OF IfcCartesianPoint;
	CurveForm : IfcBSplineCurveForm;
	ClosedCurve : IfcLogical;
	SelfIntersect : IfcLogical;
 DERIVE
	 UpperIndexOnControlPoints : IfcInteger := (SIZEOF(ControlPointsList) - 1);
	 ControlPoints : ARRAY [0:UpperIndexOnControlPoints] OF IfcCartesianPoint := IfcListToArray(ControlPointsList,0,UpperIndexOnControlPoints);
 WHERE
	SameDim : SIZEOF(QUERY(Temp <* ControlPointsList |
  Temp.Dim <> ControlPointsList[1].Dim))
= 0;
END_ENTITY;`,
	"IFCBSPLINECURVEWITHKNOTS": `ENTITY IfcBSplineCurveWithKnots
 SUPERTYPE OF (ONEOF
	(IfcRationalBSplineCurveWithKnots))
 SUBTYPE OF (IfcBSplineCurve);
	KnotMultiplicities : LIST [2:?] OF IfcInteger;
	Knots : LIST [2:?] OF IfcParameterValue;
	KnotSpec : IfcKnotType;
 DERIVE
	 UpperIndexOnKnots : IfcInteger := SIZEOF(Knots);
 WHERE
	ConsistentBSpline : IfcConstraintsParamBSpline(Degree, UpperIndexOnKnots,
UpperIndexOnControlPoints, KnotMultiplicities, Knots);
	CorrespondingKnotLists : SIZEOF(KnotMultiplicities) = UpperIndexOnKnots;
END_ENTITY;`,
	"IFCBSPLINESURFACE": `ENTITY IfcBSplineSurface
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBSplineSurfaceWithKnots))
 SUBTYPE OF (IfcBoundedSurface);
	UDegree : IfcInteger;
	VDegree : IfcInteger;
	ControlPointsList : LIST [2:?] OF LIST [2:?] OF IfcCartesianPoint;
	SurfaceForm : IfcBSplineSurfaceForm;
	UClosed : IfcLogical;
	VClosed : IfcLogical;
	SelfIntersect : IfcLogical;
 DERIVE
	 UUpper : IfcInteger := SIZEOF(ControlPointsList) - 1;
	 VUpper : IfcInteger := SIZEOF(ControlPointsList[1]) - 1;
	 ControlPoints : ARRAY [0:UUpper] OF ARRAY [0:VUpper] OF IfcCartesianPoint := IfcMakeArrayOfArray(ControlPointsList,
0,UUpper,0,VUpper);
END_ENTITY;`,
	"IFCBSPLINESURFACEWITHKNOTS": `ENTITY IfcBSplineSurfaceWithKnots
 SUPERTYPE OF (ONEOF
	(IfcRationalBSplineSurfaceWithKnots))
 SUBTYPE OF (IfcBSplineSurface);
	UMultiplicities : LIST [2:?] OF IfcInteger;
	VMultiplicities : LIST [2:?] OF IfcInteger;
	UKnots : LIST [2:?] OF IfcParameterValue;
	VKnots : LIST [2:?] OF IfcParameterValue;
	KnotSpec : IfcKnotType;
 DERIVE
	 KnotVUpper : IfcInteger := SIZEOF(VKnots);
	 KnotUUpper : IfcInteger := SIZEOF(UKnots);
 WHERE
	CorrespondingULists : SIZEOF(UMultiplicities) = KnotUUpper;
	CorrespondingVLists : SIZEOF(VMultiplicities) = KnotVUpper;
	UDirectionConstraints : IfcConstraintsParamBSpline (
  SELF\IfcBSplineSurface.UDegree, KnotUUpper, 
  SELF\IfcBSplineSurface.UUpper, UMultiplicities, UKnots);
	VDirectionConstraints : IfcConstraintsParamBSpline (
  SELF\IfcBSplineSurface.VDegree, KnotVUpper, 
  SELF\IfcBSplineSurface.VUpper, VMultiplicities, VKnots);
END_ENTITY;`,
	"IFCBUILDING": `ENTITY IfcBuilding
 SUBTYPE OF (IfcFacility);
	ElevationOfRefHeight : OPTIONAL IfcLengthMeasure;
	ElevationOfTerrain : OPTIONAL IfcLengthMeasure;
	BuildingAddress : OPTIONAL IfcPostalAddress;
END_ENTITY;`,
	"IFCBUILDINGELEMENTPART": `ENTITY IfcBuildingElementPart
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcBuildingElementPartTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBuildingElementPartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuildingElementPartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBUILDINGELEMENTPARTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCBUILDINGELEMENTPARTTYPE": `ENTITY IfcBuildingElementPartType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcBuildingElementPartTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBuildingElementPartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuildingElementPartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCBUILDINGELEMENTPROXY": `ENTITY IfcBuildingElementProxy
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcBuildingElementProxyTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBuildingElementProxyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuildingElementProxyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBUILDINGELEMENTPROXYTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
	HasObjectName : EXISTS(SELF\IfcRoot.Name);
END_ENTITY;`,
	"IFCBUILDINGELEMENTPROXYTYPE": `ENTITY IfcBuildingElementProxyType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcBuildingElementProxyTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBuildingElementProxyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuildingElementProxyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCBUILDINGSTOREY": `ENTITY IfcBuildingStorey
 SUBTYPE OF (IfcSpatialStructureElement);
	Elevation : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCBUILDINGSYSTEM": `ENTITY IfcBuildingSystem
 SUBTYPE OF (IfcSystem);
	PredefinedType : OPTIONAL IfcBuildingSystemTypeEnum;
	LongName : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBuildingSystemTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuildingSystemTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCBUILTELEMENT": `ENTITY IfcBuiltElement
 SUPERTYPE OF (ONEOF
	(IfcBeam
	,IfcBearing
	,IfcBuildingElementProxy
	,IfcChimney
	,IfcColumn
	,IfcCourse
	,IfcCovering
	,IfcCurtainWall
	,IfcDeepFoundation
	,IfcDoor
	,IfcEarthworksElement
	,IfcFooting
	,IfcKerb
	,IfcMember
	,IfcMooringDevice
	,IfcNavigationElement
	,IfcPavement
	,IfcPlate
	,IfcRail
	,IfcRailing
	,IfcRamp
	,IfcRampFlight
	,IfcRoof
	,IfcShadingDevice
	,IfcSlab
	,IfcStair
	,IfcStairFlight
	,IfcTrackElement
	,IfcWall
	,IfcWindow))
 SUBTYPE OF (IfcElement);
 WHERE
	MaxOneMaterialAssociation : SIZEOF (QUERY(temp <* SELF\IfcObjectDefinition.HasAssociations |
  'IFC4X3_DEV_73740fe4.IFCRELASSOCIATESMATERIAL' IN TYPEOF(temp)
  )) <= 1;
END_ENTITY;`,
	"IFCBUILTELEMENTTYPE": `ENTITY IfcBuiltElementType
 SUPERTYPE OF (ONEOF
	(IfcBeamType
	,IfcBearingType
	,IfcBuildingElementProxyType
	,IfcChimneyType
	,IfcColumnType
	,IfcCourseType
	,IfcCoveringType
	,IfcCurtainWallType
	,IfcDeepFoundationType
	,IfcDoorType
	,IfcFootingType
	,IfcKerbType
	,IfcMemberType
	,IfcMooringDeviceType
	,IfcNavigationElementType
	,IfcPavementType
	,IfcPlateType
	,IfcRailType
	,IfcRailingType
	,IfcRampFlightType
	,IfcRampType
	,IfcRoofType
	,IfcShadingDeviceType
	,IfcSlabType
	,IfcStairFlightType
	,IfcStairType
	,IfcTrackElementType
	,IfcWallType
	,IfcWindowType))
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCBUILTSYSTEM": `ENTITY IfcBuiltSystem
 SUBTYPE OF (IfcSystem);
	PredefinedType : OPTIONAL IfcBuiltSystemTypeEnum;
	LongName : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBuiltSystemTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBuiltSystemTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCBURNER": `ENTITY IfcBurner
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcBurnerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcBurnerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBurnerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCBURNERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCBURNERTYPE": `ENTITY IfcBurnerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcBurnerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcBurnerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcBurnerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCABLECARRIERFITTING": `ENTITY IfcCableCarrierFitting
 SUBTYPE OF (IfcFlowFitting);
	PredefinedType : OPTIONAL IfcCableCarrierFittingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCableCarrierFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableCarrierFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCABLECARRIERFITTINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCABLECARRIERFITTINGTYPE": `ENTITY IfcCableCarrierFittingType
 SUBTYPE OF (IfcFlowFittingType);
	PredefinedType : IfcCableCarrierFittingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCableCarrierFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableCarrierFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCABLECARRIERSEGMENT": `ENTITY IfcCableCarrierSegment
 SUBTYPE OF (IfcFlowSegment);
	PredefinedType : OPTIONAL IfcCableCarrierSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCableCarrierSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableCarrierSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCABLECARRIERSEGMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCABLECARRIERSEGMENTTYPE": `ENTITY IfcCableCarrierSegmentType
 SUBTYPE OF (IfcFlowSegmentType);
	PredefinedType : IfcCableCarrierSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCableCarrierSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableCarrierSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCABLEFITTING": `ENTITY IfcCableFitting
 SUBTYPE OF (IfcFlowFitting);
	PredefinedType : OPTIONAL IfcCableFittingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCableFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCABLEFITTINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCABLEFITTINGTYPE": `ENTITY IfcCableFittingType
 SUBTYPE OF (IfcFlowFittingType);
	PredefinedType : IfcCableFittingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCableFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCABLESEGMENT": `ENTITY IfcCableSegment
 SUBTYPE OF (IfcFlowSegment);
	PredefinedType : OPTIONAL IfcCableSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCableSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCABLESEGMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCABLESEGMENTTYPE": `ENTITY IfcCableSegmentType
 SUBTYPE OF (IfcFlowSegmentType);
	PredefinedType : IfcCableSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCableSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCableSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCAISSONFOUNDATION": `ENTITY IfcCaissonFoundation
 SUBTYPE OF (IfcDeepFoundation);
	PredefinedType : OPTIONAL IfcCaissonFoundationTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCaissonFoundationTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCaissonFoundationTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCAISSONFOUNDATIONTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCAISSONFOUNDATIONTYPE": `ENTITY IfcCaissonFoundationType
 SUBTYPE OF (IfcDeepFoundationType);
	PredefinedType : IfcCaissonFoundationTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCaissonFoundationTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCaissonFoundationTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCARTESIANPOINT": `ENTITY IfcCartesianPoint
 SUBTYPE OF (IfcPoint);
	Coordinates : LIST [1:3] OF IfcLengthMeasure;
 WHERE
	CP2Dor3D : HIINDEX(Coordinates) >= 2;
END_ENTITY;`,
	"IFCCARTESIANPOINTLIST": `ENTITY IfcCartesianPointList
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCartesianPointList2D
	,IfcCartesianPointList3D))
 SUBTYPE OF (IfcGeometricRepresentationItem);
 DERIVE
	 Dim : IfcDimensionCount := IfcPointListDim(SELF);
END_ENTITY;`,
	"IFCCARTESIANPOINTLIST2D": `ENTITY IfcCartesianPointList2D
 SUBTYPE OF (IfcCartesianPointList);
	CoordList : LIST [1:?] OF LIST [2:2] OF IfcLengthMeasure;
	TagList : OPTIONAL LIST [1:?] OF IfcLabel;
END_ENTITY;`,
	"IFCCARTESIANPOINTLIST3D": `ENTITY IfcCartesianPointList3D
 SUBTYPE OF (IfcCartesianPointList);
	CoordList : LIST [1:?] OF LIST [3:3] OF IfcLengthMeasure;
	TagList : OPTIONAL LIST [1:?] OF IfcLabel;
END_ENTITY;`,
	"IFCCARTESIANTRANSFORMATIONOPERATOR": `ENTITY IfcCartesianTransformationOperator
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCartesianTransformationOperator2D
	,IfcCartesianTransformationOperator3D))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Axis1 : OPTIONAL IfcDirection;
	Axis2 : OPTIONAL IfcDirection;
	LocalOrigin : IfcCartesianPoint;
	Scale : OPTIONAL IfcReal;
 DERIVE
	 Scl : IfcReal := NVL(Scale, 1.0);
	 Dim : IfcDimensionCount := LocalOrigin.Dim;
 WHERE
	ScaleGreaterZero : Scl > 0.0;
END_ENTITY;`,
	"IFCCARTESIANTRANSFORMATIONOPERATOR2D": `ENTITY IfcCartesianTransformationOperator2D
 SUPERTYPE OF (ONEOF
	(IfcCartesianTransformationOperator2DnonUniform))
 SUBTYPE OF (IfcCartesianTransformationOperator);
 DERIVE
	 U : LIST [2:2] OF IfcDirection := IfcBaseAxis(2,SELF\IfcCartesianTransformationOperator.Axis1,
SELF\IfcCartesianTransformationOperator.Axis2,?);
 WHERE
	Axis1Is2D : NOT(EXISTS(SELF\IfcCartesianTransformationOperator.Axis1)) OR 
(SELF\IfcCartesianTransformationOperator.Axis1.Dim = 2);
	Axis2Is2D : NOT(EXISTS(SELF\IfcCartesianTransformationOperator.Axis2)) OR 
(SELF\IfcCartesianTransformationOperator.Axis2.Dim = 2);
	DimEqual2 : SELF\IfcCartesianTransformationOperator.Dim = 2;
END_ENTITY;`,
	"IFCCARTESIANTRANSFORMATIONOPERATOR2DNONUNIFORM": `ENTITY IfcCartesianTransformationOperator2DnonUniform
 SUBTYPE OF (IfcCartesianTransformationOperator2D);
	Scale2 : OPTIONAL IfcReal;
 DERIVE
	 Scl2 : IfcReal := NVL(Scale2, SELF\IfcCartesianTransformationOperator.Scl);
 WHERE
	Scale2GreaterZero : Scl2 > 0.0;
END_ENTITY;`,
	"IFCCARTESIANTRANSFORMATIONOPERATOR3D": `ENTITY IfcCartesianTransformationOperator3D
 SUPERTYPE OF (ONEOF
	(IfcCartesianTransformationOperator3DnonUniform))
 SUBTYPE OF (IfcCartesianTransformationOperator);
	Axis3 : OPTIONAL IfcDirection;
 DERIVE
	 U : LIST [3:3] OF IfcDirection := IfcBaseAxis(3,SELF\IfcCartesianTransformationOperator.Axis1,
SELF\IfcCartesianTransformationOperator.Axis2,Axis3);
 WHERE
	Axis1Is3D : NOT(EXISTS(SELF\IfcCartesianTransformationOperator.Axis1)) OR 
(SELF\IfcCartesianTransformationOperator.Axis1.Dim = 3);
	Axis2Is3D : NOT(EXISTS(SELF\IfcCartesianTransformationOperator.Axis2)) OR 
(SELF\IfcCartesianTransformationOperator.Axis2.Dim = 3);
	Axis3Is3D : NOT(EXISTS(Axis3)) OR (Axis3.Dim = 3);
	DimIs3D : SELF\IfcCartesianTransformationOperator.Dim = 3;
END_ENTITY;`,
	"IFCCARTESIANTRANSFORMATIONOPERATOR3DNONUNIFORM": `ENTITY IfcCartesianTransformationOperator3DnonUniform
 SUBTYPE OF (IfcCartesianTransformationOperator3D);
	Scale2 : OPTIONAL IfcReal;
	Scale3 : OPTIONAL IfcReal;
 DERIVE
	 Scl2 : IfcReal := NVL(Scale2, SELF\IfcCartesianTransformationOperator.Scl);
	 Scl3 : IfcReal := NVL(Scale3, SELF\IfcCartesianTransformationOperator.Scl);
 WHERE
	Scale2GreaterZero : Scl2 > 0.0;
	Scale3GreaterZero : Scl3 > 0.0;
END_ENTITY;`,
	"IFCCENTERLINEPROFILEDEF": `ENTITY IfcCenterLineProfileDef
 SUBTYPE OF (IfcArbitraryOpenProfileDef);
	Thickness : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCCHILLER": `ENTITY IfcChiller
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcChillerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcChillerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcChillerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCHILLERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCHILLERTYPE": `ENTITY IfcChillerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcChillerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcChillerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcChillerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCHIMNEY": `ENTITY IfcChimney
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcChimneyTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcChimneyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcChimneyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCHIMNEYTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCHIMNEYTYPE": `ENTITY IfcChimneyType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcChimneyTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcChimneyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcChimneyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCIRCLE": `ENTITY IfcCircle
 SUBTYPE OF (IfcConic);
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCCIRCLEHOLLOWPROFILEDEF": `ENTITY IfcCircleHollowProfileDef
 SUBTYPE OF (IfcCircleProfileDef);
	WallThickness : IfcPositiveLengthMeasure;
 WHERE
	WR1 : WallThickness < SELF\IfcCircleProfileDef.Radius;
END_ENTITY;`,
	"IFCCIRCLEPROFILEDEF": `ENTITY IfcCircleProfileDef
 SUPERTYPE OF (ONEOF
	(IfcCircleHollowProfileDef))
 SUBTYPE OF (IfcParameterizedProfileDef);
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCCIVILELEMENT": `ENTITY IfcCivilElement
 SUBTYPE OF (IfcElement);
END_ENTITY;`,
	"IFCCIVILELEMENTTYPE": `ENTITY IfcCivilElementType
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCCLASSIFICATION": `ENTITY IfcClassification
 SUBTYPE OF (IfcExternalInformation);
	Source : OPTIONAL IfcLabel;
	Edition : OPTIONAL IfcLabel;
	EditionDate : OPTIONAL IfcDate;
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	Specification : OPTIONAL IfcURIReference;
	ReferenceTokens : OPTIONAL LIST [1:?] OF IfcIdentifier;
 INVERSE
	ClassificationForObjects : SET [0:?] OF IfcRelAssociatesClassification FOR RelatingClassification;
	HasReferences : SET [0:?] OF IfcClassificationReference FOR ReferencedSource;
END_ENTITY;`,
	"IFCCLASSIFICATIONREFERENCE": `ENTITY IfcClassificationReference
 SUBTYPE OF (IfcExternalReference);
	ReferencedSource : OPTIONAL IfcClassificationReferenceSelect;
	Description : OPTIONAL IfcText;
	Sort : OPTIONAL IfcIdentifier;
 INVERSE
	ClassificationRefForObjects : SET [0:?] OF IfcRelAssociatesClassification FOR RelatingClassification;
	HasReferences : SET [0:?] OF IfcClassificationReference FOR ReferencedSource;
END_ENTITY;`,
	"IFCCLOSEDSHELL": `ENTITY IfcClosedShell
 SUBTYPE OF (IfcConnectedFaceSet);
END_ENTITY;`,
	"IFCCLOTHOID": `ENTITY IfcClothoid
 SUBTYPE OF (IfcSpiral);
	ClothoidConstant : IfcLengthMeasure;
END_ENTITY;`,
	"IFCCOIL": `ENTITY IfcCoil
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcCoilTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCoilTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoilTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOILTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOILTYPE": `ENTITY IfcCoilType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcCoilTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCoilTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoilTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOLOURRGB": `ENTITY IfcColourRgb
 SUBTYPE OF (IfcColourSpecification);
	Red : IfcNormalisedRatioMeasure;
	Green : IfcNormalisedRatioMeasure;
	Blue : IfcNormalisedRatioMeasure;
END_ENTITY;`,
	"IFCCOLOURRGBLIST": `ENTITY IfcColourRgbList
 SUBTYPE OF (IfcPresentationItem);
	ColourList : LIST [1:?] OF LIST [3:3] OF IfcNormalisedRatioMeasure;
END_ENTITY;`,
	"IFCCOLOURSPECIFICATION": `ENTITY IfcColourSpecification
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcColourRgb))
 SUBTYPE OF (IfcPresentationItem);
	Name : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCCOLUMN": `ENTITY IfcColumn
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcColumnTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcColumnTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcColumnTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOLUMNTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOLUMNTYPE": `ENTITY IfcColumnType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcColumnTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcColumnTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcColumnTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOMMUNICATIONSAPPLIANCE": `ENTITY IfcCommunicationsAppliance
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcCommunicationsApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCommunicationsApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCommunicationsApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOMMUNICATIONSAPPLIANCETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOMMUNICATIONSAPPLIANCETYPE": `ENTITY IfcCommunicationsApplianceType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcCommunicationsApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCommunicationsApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCommunicationsApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOMPLEXPROPERTY": `ENTITY IfcComplexProperty
 SUBTYPE OF (IfcProperty);
	UsageName : IfcIdentifier;
	HasProperties : SET [1:?] OF IfcProperty;
 WHERE
	WR21 : SIZEOF(QUERY(temp <* HasProperties | SELF :=: temp)) = 0;
	WR22 : IfcUniquePropertyName(HasProperties);
END_ENTITY;`,
	"IFCCOMPLEXPROPERTYTEMPLATE": `ENTITY IfcComplexPropertyTemplate
 SUBTYPE OF (IfcPropertyTemplate);
	UsageName : OPTIONAL IfcLabel;
	TemplateType : OPTIONAL IfcComplexPropertyTemplateTypeEnum;
	HasPropertyTemplates : OPTIONAL SET [1:?] OF IfcPropertyTemplate;
 WHERE
	NoSelfReference : SIZEOF(QUERY(temp <* HasPropertyTemplates | SELF :=: temp)) = 0;
	UniquePropertyNames : IfcUniquePropertyTemplateNames(HasPropertyTemplates);
END_ENTITY;`,
	"IFCCOMPOSITECURVE": `ENTITY IfcCompositeCurve
 SUPERTYPE OF (ONEOF
	(IfcCompositeCurveOnSurface
	,IfcGradientCurve
	,IfcSegmentedReferenceCurve))
 SUBTYPE OF (IfcBoundedCurve);
	Segments : LIST [1:?] OF IfcSegment;
	SelfIntersect : IfcLogical;
 DERIVE
	 NSegments : IfcInteger := SIZEOF(Segments);
	 ClosedCurve : IfcLogical := Segments[NSegments].Transition <> Discontinuous;
 WHERE
	CurveContinuous : ((NOT ClosedCurve) AND (SIZEOF(QUERY(Temp <* Segments | Temp.Transition = Discontinuous)) = 1)) OR ((ClosedCurve) AND (SIZEOF(QUERY(Temp <* Segments | Temp.Transition = Discontinuous)) = 0));
	SameDim : SIZEOF( QUERY( Temp <* Segments | Temp.Dim <> Segments[1].Dim)) = 0;
END_ENTITY;`,
	"IFCCOMPOSITECURVEONSURFACE": `ENTITY IfcCompositeCurveOnSurface
 SUPERTYPE OF (ONEOF
	(IfcBoundaryCurve))
 SUBTYPE OF (IfcCompositeCurve);
 DERIVE
	 BasisSurface : SET [0:1] OF IfcSurface := IfcGetBasisSurface(SELF);
 WHERE
	SameSurface : SIZEOF(BasisSurface) > 0;
END_ENTITY;`,
	"IFCCOMPOSITECURVESEGMENT": `ENTITY IfcCompositeCurveSegment
 SUPERTYPE OF (ONEOF
	(IfcReparametrisedCompositeCurveSegment))
 SUBTYPE OF (IfcSegment);
	SameSense : IfcBoolean;
	ParentCurve : IfcCurve;
 WHERE
	ParentIsBoundedCurve : ('IFC4X3_DEV_73740fe4.IFCBOUNDEDCURVE' IN TYPEOF(ParentCurve));
END_ENTITY;`,
	"IFCCOMPOSITEPROFILEDEF": `ENTITY IfcCompositeProfileDef
 SUBTYPE OF (IfcProfileDef);
	Profiles : SET [2:?] OF IfcProfileDef;
	Label : OPTIONAL IfcLabel;
 WHERE
	InvariantProfileType : SIZEOF(QUERY(temp <* Profiles | temp.ProfileType <> Profiles[1].ProfileType)) = 0;
	NoRecursion : SIZEOF(QUERY(temp <* Profiles | 'IFC4X3_DEV_73740fe4.IFCCOMPOSITEPROFILEDEF' IN TYPEOF(temp))) = 0;
END_ENTITY;`,
	"IFCCOMPRESSOR": `ENTITY IfcCompressor
 SUBTYPE OF (IfcFlowMovingDevice);
	PredefinedType : OPTIONAL IfcCompressorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCompressorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCompressorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOMPRESSORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOMPRESSORTYPE": `ENTITY IfcCompressorType
 SUBTYPE OF (IfcFlowMovingDeviceType);
	PredefinedType : IfcCompressorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCompressorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCompressorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCONDENSER": `ENTITY IfcCondenser
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcCondenserTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCondenserTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCondenserTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCONDENSERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCONDENSERTYPE": `ENTITY IfcCondenserType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcCondenserTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCondenserTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCondenserTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCONIC": `ENTITY IfcConic
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCircle
	,IfcEllipse))
 SUBTYPE OF (IfcCurve);
	Position : IfcAxis2Placement;
END_ENTITY;`,
	"IFCCONNECTEDFACESET": `ENTITY IfcConnectedFaceSet
 SUPERTYPE OF (ONEOF
	(IfcClosedShell
	,IfcOpenShell))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
	CfsFaces : SET [1:?] OF IfcFace;
END_ENTITY;`,
	"IFCCONNECTIONCURVEGEOMETRY": `ENTITY IfcConnectionCurveGeometry
 SUBTYPE OF (IfcConnectionGeometry);
	CurveOnRelatingElement : IfcCurveOrEdgeCurve;
	CurveOnRelatedElement : OPTIONAL IfcCurveOrEdgeCurve;
END_ENTITY;`,
	"IFCCONNECTIONGEOMETRY": `ENTITY IfcConnectionGeometry
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConnectionCurveGeometry
	,IfcConnectionPointGeometry
	,IfcConnectionSurfaceGeometry
	,IfcConnectionVolumeGeometry));
END_ENTITY;`,
	"IFCCONNECTIONPOINTECCENTRICITY": `ENTITY IfcConnectionPointEccentricity
 SUBTYPE OF (IfcConnectionPointGeometry);
	EccentricityInX : OPTIONAL IfcLengthMeasure;
	EccentricityInY : OPTIONAL IfcLengthMeasure;
	EccentricityInZ : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCCONNECTIONPOINTGEOMETRY": `ENTITY IfcConnectionPointGeometry
 SUPERTYPE OF (ONEOF
	(IfcConnectionPointEccentricity))
 SUBTYPE OF (IfcConnectionGeometry);
	PointOnRelatingElement : IfcPointOrVertexPoint;
	PointOnRelatedElement : OPTIONAL IfcPointOrVertexPoint;
END_ENTITY;`,
	"IFCCONNECTIONSURFACEGEOMETRY": `ENTITY IfcConnectionSurfaceGeometry
 SUBTYPE OF (IfcConnectionGeometry);
	SurfaceOnRelatingElement : IfcSurfaceOrFaceSurface;
	SurfaceOnRelatedElement : OPTIONAL IfcSurfaceOrFaceSurface;
END_ENTITY;`,
	"IFCCONNECTIONVOLUMEGEOMETRY": `ENTITY IfcConnectionVolumeGeometry
 SUBTYPE OF (IfcConnectionGeometry);
	VolumeOnRelatingElement : IfcSolidOrShell;
	VolumeOnRelatedElement : OPTIONAL IfcSolidOrShell;
END_ENTITY;`,
	"IFCCONSTRAINT": `ENTITY IfcConstraint
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMetric
	,IfcObjective));
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	ConstraintGrade : IfcConstraintEnum;
	ConstraintSource : OPTIONAL IfcLabel;
	CreatingActor : OPTIONAL IfcActorSelect;
	CreationTime : OPTIONAL IfcDateTime;
	UserDefinedGrade : OPTIONAL IfcLabel;
 INVERSE
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
	PropertiesForConstraint : SET [0:?] OF IfcResourceConstraintRelationship FOR RelatingConstraint;
 WHERE
	WR11 : (ConstraintGrade <> IfcConstraintEnum.USERDEFINED) OR
((ConstraintGrade = IfcConstraintEnum.USERDEFINED) AND EXISTS(SELF\IfcConstraint.UserDefinedGrade));
END_ENTITY;`,
	"IFCCONSTRUCTIONEQUIPMENTRESOURCE": `ENTITY IfcConstructionEquipmentResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcConstructionEquipmentResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
(PredefinedType <> IfcConstructionEquipmentResourceTypeEnum.USERDEFINED) OR 
((PredefinedType = IfcConstructionEquipmentResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCCONSTRUCTIONEQUIPMENTRESOURCETYPE": `ENTITY IfcConstructionEquipmentResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcConstructionEquipmentResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcConstructionEquipmentResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcConstructionEquipmentResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCCONSTRUCTIONMATERIALRESOURCE": `ENTITY IfcConstructionMaterialResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcConstructionMaterialResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
 (PredefinedType <> IfcConstructionMaterialResourceTypeEnum.USERDEFINED) OR 
 ((PredefinedType = IfcConstructionMaterialResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCCONSTRUCTIONMATERIALRESOURCETYPE": `ENTITY IfcConstructionMaterialResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcConstructionMaterialResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcConstructionMaterialResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcConstructionMaterialResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCCONSTRUCTIONPRODUCTRESOURCE": `ENTITY IfcConstructionProductResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcConstructionProductResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
 (PredefinedType <> IfcConstructionProductResourceTypeEnum.USERDEFINED) OR 
 ((PredefinedType = IfcConstructionProductResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCCONSTRUCTIONPRODUCTRESOURCETYPE": `ENTITY IfcConstructionProductResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcConstructionProductResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcConstructionProductResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcConstructionProductResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCCONSTRUCTIONRESOURCE": `ENTITY IfcConstructionResource
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConstructionEquipmentResource
	,IfcConstructionMaterialResource
	,IfcConstructionProductResource
	,IfcCrewResource
	,IfcLaborResource
	,IfcSubContractResource))
 SUBTYPE OF (IfcResource);
	Usage : OPTIONAL IfcResourceTime;
	BaseCosts : OPTIONAL LIST [1:?] OF IfcAppliedValue;
	BaseQuantity : OPTIONAL IfcPhysicalQuantity;
END_ENTITY;`,
	"IFCCONSTRUCTIONRESOURCETYPE": `ENTITY IfcConstructionResourceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConstructionEquipmentResourceType
	,IfcConstructionMaterialResourceType
	,IfcConstructionProductResourceType
	,IfcCrewResourceType
	,IfcLaborResourceType
	,IfcSubContractResourceType))
 SUBTYPE OF (IfcTypeResource);
	BaseCosts : OPTIONAL LIST [1:?] OF IfcAppliedValue;
	BaseQuantity : OPTIONAL IfcPhysicalQuantity;
END_ENTITY;`,
	"IFCCONTEXT": `ENTITY IfcContext
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcProject
	,IfcProjectLibrary))
 SUBTYPE OF (IfcObjectDefinition);
	ObjectType : OPTIONAL IfcLabel;
	LongName : OPTIONAL IfcLabel;
	Phase : OPTIONAL IfcLabel;
	RepresentationContexts : OPTIONAL SET [1:?] OF IfcRepresentationContext;
	UnitsInContext : OPTIONAL IfcUnitAssignment;
 INVERSE
	IsDefinedBy : SET [0:?] OF IfcRelDefinesByProperties FOR RelatedObjects;
	Declares : SET [0:?] OF IfcRelDeclares FOR RelatingContext;
END_ENTITY;`,
	"IFCCONTEXTDEPENDENTUNIT": `ENTITY IfcContextDependentUnit
 SUBTYPE OF (IfcNamedUnit);
	Name : IfcLabel;
 INVERSE
	HasExternalReference : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCCONTROL": `ENTITY IfcControl
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcActionRequest
	,IfcCostItem
	,IfcCostSchedule
	,IfcPerformanceHistory
	,IfcPermit
	,IfcProjectOrder
	,IfcWorkCalendar
	,IfcWorkControl))
 SUBTYPE OF (IfcObject);
	Identification : OPTIONAL IfcIdentifier;
 INVERSE
	Controls : SET [0:?] OF IfcRelAssignsToControl FOR RelatingControl;
END_ENTITY;`,
	"IFCCONTROLLER": `ENTITY IfcController
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcControllerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcControllerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcControllerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCONTROLLERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCONTROLLERTYPE": `ENTITY IfcControllerType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcControllerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcControllerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcControllerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCONVERSIONBASEDUNIT": `ENTITY IfcConversionBasedUnit
 SUPERTYPE OF (ONEOF
	(IfcConversionBasedUnitWithOffset))
 SUBTYPE OF (IfcNamedUnit);
	Name : IfcLabel;
	ConversionFactor : IfcMeasureWithUnit;
 INVERSE
	HasExternalReference : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCCONVERSIONBASEDUNITWITHOFFSET": `ENTITY IfcConversionBasedUnitWithOffset
 SUBTYPE OF (IfcConversionBasedUnit);
	ConversionOffset : IfcReal;
END_ENTITY;`,
	"IFCCONVEYORSEGMENT": `ENTITY IfcConveyorSegment
 SUBTYPE OF (IfcFlowSegment);
	PredefinedType : OPTIONAL IfcConveyorSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcConveyorSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcConveyorSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCONVEYORSEGMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCONVEYORSEGMENTTYPE": `ENTITY IfcConveyorSegmentType
 SUBTYPE OF (IfcFlowSegmentType);
	PredefinedType : IfcConveyorSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcConveyorSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcConveyorSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOOLEDBEAM": `ENTITY IfcCooledBeam
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcCooledBeamTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCooledBeamTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCooledBeamTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOOLEDBEAMTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOOLEDBEAMTYPE": `ENTITY IfcCooledBeamType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcCooledBeamTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCooledBeamTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCooledBeamTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOOLINGTOWER": `ENTITY IfcCoolingTower
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcCoolingTowerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCoolingTowerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoolingTowerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOOLINGTOWERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOOLINGTOWERTYPE": `ENTITY IfcCoolingTowerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcCoolingTowerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCoolingTowerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoolingTowerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOORDINATEOPERATION": `ENTITY IfcCoordinateOperation
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMapConversion
	,IfcRigidOperation));
	SourceCRS : IfcCoordinateReferenceSystemSelect;
	TargetCRS : IfcCoordinateReferenceSystem;
END_ENTITY;`,
	"IFCCOORDINATEREFERENCESYSTEM": `ENTITY IfcCoordinateReferenceSystem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGeographicCRS
	,IfcProjectedCRS));
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	GeodeticDatum : OPTIONAL IfcIdentifier;
 INVERSE
	HasCoordinateOperation : SET [0:1] OF IfcCoordinateOperation FOR SourceCRS;
	WellKnownText : SET [0:1] OF IfcWellKnownText FOR CoordinateReferenceSystem;
 WHERE
	NameOrWKT : (HIINDEX(WellKnownText) = 1) OR EXISTS(Name);
END_ENTITY;`,
	"IFCCOSINESPIRAL": `ENTITY IfcCosineSpiral
 SUBTYPE OF (IfcSpiral);
	CosineTerm : IfcLengthMeasure;
	ConstantTerm : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCCOSTITEM": `ENTITY IfcCostItem
 SUBTYPE OF (IfcControl);
	PredefinedType : OPTIONAL IfcCostItemTypeEnum;
	CostValues : OPTIONAL LIST [1:?] OF IfcCostValue;
	CostQuantities : OPTIONAL LIST [1:?] OF IfcPhysicalQuantity;
END_ENTITY;`,
	"IFCCOSTSCHEDULE": `ENTITY IfcCostSchedule
 SUBTYPE OF (IfcControl);
	PredefinedType : OPTIONAL IfcCostScheduleTypeEnum;
	Status : OPTIONAL IfcLabel;
	SubmittedOn : OPTIONAL IfcDateTime;
	UpdateDate : OPTIONAL IfcDateTime;
END_ENTITY;`,
	"IFCCOSTVALUE": `ENTITY IfcCostValue
 SUBTYPE OF (IfcAppliedValue);
END_ENTITY;`,
	"IFCCOURSE": `ENTITY IfcCourse
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcCourseTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCourseTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCourseTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOURSETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOURSETYPE": `ENTITY IfcCourseType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcCourseTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCourseTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCourseTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCOVERING": `ENTITY IfcCovering
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcCoveringTypeEnum;
 INVERSE
	CoversSpaces : SET [0:1] OF IfcRelCoversSpaces FOR RelatedCoverings;
	CoversElements : SET [0:1] OF IfcRelCoversBldgElements FOR RelatedCoverings;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCoveringTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoveringTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCOVERINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCOVERINGTYPE": `ENTITY IfcCoveringType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcCoveringTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCoveringTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCoveringTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCREWRESOURCE": `ENTITY IfcCrewResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcCrewResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
 (PredefinedType <> IfcCrewResourceTypeEnum.USERDEFINED) OR 
 ((PredefinedType = IfcCrewResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCCREWRESOURCETYPE": `ENTITY IfcCrewResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcCrewResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCrewResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcCrewResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCCSGPRIMITIVE3D": `ENTITY IfcCsgPrimitive3D
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBlock
	,IfcRectangularPyramid
	,IfcRightCircularCone
	,IfcRightCircularCylinder
	,IfcSphere))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Position : IfcAxis2Placement3D;
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCCSGSOLID": `ENTITY IfcCsgSolid
 SUBTYPE OF (IfcSolidModel);
	TreeRootExpression : IfcCsgSelect;
END_ENTITY;`,
	"IFCCSHAPEPROFILEDEF": `ENTITY IfcCShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	Depth : IfcPositiveLengthMeasure;
	Width : IfcPositiveLengthMeasure;
	WallThickness : IfcPositiveLengthMeasure;
	Girth : IfcPositiveLengthMeasure;
	InternalFilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
 WHERE
	ValidGirth : Girth < (Depth / 2.);
	ValidInternalFilletRadius : NOT(EXISTS(InternalFilletRadius)) OR
((InternalFilletRadius <= Width/2. - WallThickness) AND (InternalFilletRadius <= Depth/2. - WallThickness));
	ValidWallThickness : (WallThickness < Width/2.) AND (WallThickness < Depth/2.);
END_ENTITY;`,
	"IFCCURRENCYRELATIONSHIP": `ENTITY IfcCurrencyRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingMonetaryUnit : IfcMonetaryUnit;
	RelatedMonetaryUnit : IfcMonetaryUnit;
	ExchangeRate : IfcPositiveRatioMeasure;
	RateDateTime : OPTIONAL IfcDateTime;
	RateSource : OPTIONAL IfcLibraryInformation;
END_ENTITY;`,
	"IFCCURTAINWALL": `ENTITY IfcCurtainWall
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcCurtainWallTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcCurtainWallTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCurtainWallTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCCURTAINWALLTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCCURTAINWALLTYPE": `ENTITY IfcCurtainWallType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcCurtainWallTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcCurtainWallTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcCurtainWallTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCCURVE": `ENTITY IfcCurve
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBoundedCurve
	,IfcConic
	,IfcLine
	,IfcOffsetCurve
	,IfcPcurve
	,IfcPolynomialCurve
	,IfcSpiral
	,IfcSurfaceCurve))
 SUBTYPE OF (IfcGeometricRepresentationItem);
 DERIVE
	 Dim : IfcDimensionCount := IfcCurveDim(SELF);
END_ENTITY;`,
	"IFCCURVEBOUNDEDPLANE": `ENTITY IfcCurveBoundedPlane
 SUBTYPE OF (IfcBoundedSurface);
	BasisSurface : IfcPlane;
	OuterBoundary : IfcCurve;
	InnerBoundaries : SET [0:?] OF IfcCurve;
END_ENTITY;`,
	"IFCCURVEBOUNDEDSURFACE": `ENTITY IfcCurveBoundedSurface
 SUBTYPE OF (IfcBoundedSurface);
	BasisSurface : IfcSurface;
	Boundaries : SET [1:?] OF IfcBoundaryCurve;
	ImplicitOuter : IfcBoolean;
END_ENTITY;`,
	"IFCCURVESEGMENT": `ENTITY IfcCurveSegment
 SUBTYPE OF (IfcSegment);
	Placement : IfcPlacement;
	SegmentStart : IfcCurveMeasureSelect;
	SegmentLength : IfcCurveMeasureSelect;
	ParentCurve : IfcCurve;
END_ENTITY;`,
	"IFCCURVESTYLE": `ENTITY IfcCurveStyle
 SUBTYPE OF (IfcPresentationStyle);
	CurveFont : OPTIONAL IfcCurveFontOrScaledCurveFontSelect;
	CurveWidth : OPTIONAL IfcSizeSelect;
	CurveColour : OPTIONAL IfcColour;
	ModelOrDraughting : OPTIONAL IfcBoolean;
 WHERE
	IdentifiableCurveStyle : EXISTS(CurveFont) OR EXISTS(CurveWidth) OR EXISTS(CurveColour);
	MeasureOfWidth : (NOT(EXISTS(CurveWidth))) OR
('IFC4X3_DEV_73740fe4.IFCPOSITIVELENGTHMEASURE' IN TYPEOF(CurveWidth)) OR 
 (('IFC4X3_DEV_73740fe4.IFCDESCRIPTIVEMEASURE' IN TYPEOF(CurveWidth)) AND
 (CurveWidth = 'by layer'));
END_ENTITY;`,
	"IFCCURVESTYLEFONT": `ENTITY IfcCurveStyleFont
 SUBTYPE OF (IfcPresentationItem);
	Name : OPTIONAL IfcLabel;
	PatternList : LIST [1:?] OF IfcCurveStyleFontPattern;
END_ENTITY;`,
	"IFCCURVESTYLEFONTANDSCALING": `ENTITY IfcCurveStyleFontAndScaling
 SUBTYPE OF (IfcPresentationItem);
	Name : OPTIONAL IfcLabel;
	CurveStyleFont : IfcCurveStyleFontSelect;
	CurveFontScaling : IfcPositiveRatioMeasure;
END_ENTITY;`,
	"IFCCURVESTYLEFONTPATTERN": `ENTITY IfcCurveStyleFontPattern
 SUBTYPE OF (IfcPresentationItem);
	VisibleSegmentLength : IfcLengthMeasure;
	InvisibleSegmentLength : IfcPositiveLengthMeasure;
 WHERE
	VisibleLengthGreaterEqualZero : VisibleSegmentLength >= 0.;
END_ENTITY;`,
	"IFCCYLINDRICALSURFACE": `ENTITY IfcCylindricalSurface
 SUBTYPE OF (IfcElementarySurface);
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCDAMPER": `ENTITY IfcDamper
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcDamperTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDamperTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDamperTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDAMPERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDAMPERTYPE": `ENTITY IfcDamperType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcDamperTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDamperTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDamperTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDEEPFOUNDATION": `ENTITY IfcDeepFoundation
 SUPERTYPE OF (ONEOF
	(IfcCaissonFoundation
	,IfcPile))
 SUBTYPE OF (IfcBuiltElement);
 WHERE
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDEEPFOUNDATIONTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDEEPFOUNDATIONTYPE": `ENTITY IfcDeepFoundationType
 SUPERTYPE OF (ONEOF
	(IfcCaissonFoundationType
	,IfcPileType))
 SUBTYPE OF (IfcBuiltElementType);
END_ENTITY;`,
	"IFCDERIVEDPROFILEDEF": `ENTITY IfcDerivedProfileDef
 SUPERTYPE OF (ONEOF
	(IfcMirroredProfileDef))
 SUBTYPE OF (IfcProfileDef);
	ParentProfile : IfcProfileDef;
	Operator : IfcCartesianTransformationOperator2D;
	Label : OPTIONAL IfcLabel;
 WHERE
	InvariantProfileType : SELF\IfcProfileDef.ProfileType = ParentProfile.ProfileType;
END_ENTITY;`,
	"IFCDERIVEDUNIT": `ENTITY IfcDerivedUnit;
	Elements : SET [1:?] OF IfcDerivedUnitElement;
	UnitType : IfcDerivedUnitEnum;
	UserDefinedType : OPTIONAL IfcLabel;
	Name : OPTIONAL IfcLabel;
 DERIVE
	 Dimensions : IfcDimensionalExponents := IfcDeriveDimensionalExponents(Elements);
 WHERE
	WR1 : (SIZEOF (Elements) > 1) OR ((SIZEOF (Elements) = 1) AND (Elements[1].Exponent <> 1 ));
	WR2 : (UnitType <> IfcDerivedUnitEnum.USERDEFINED) OR
((UnitType = IfcDerivedUnitEnum.USERDEFINED) AND 
 (EXISTS(SELF.UserDefinedType)));
END_ENTITY;`,
	"IFCDERIVEDUNITELEMENT": `ENTITY IfcDerivedUnitElement;
	Unit : IfcNamedUnit;
	Exponent : INTEGER;
END_ENTITY;`,
	"IFCDIMENSIONALEXPONENTS": `ENTITY IfcDimensionalExponents;
	LengthExponent : INTEGER;
	MassExponent : INTEGER;
	TimeExponent : INTEGER;
	ElectricCurrentExponent : INTEGER;
	ThermodynamicTemperatureExponent : INTEGER;
	AmountOfSubstanceExponent : INTEGER;
	LuminousIntensityExponent : INTEGER;
END_ENTITY;`,
	"IFCDIRECTION": `ENTITY IfcDirection
 SUBTYPE OF (IfcGeometricRepresentationItem);
	DirectionRatios : LIST [2:3] OF IfcReal;
 DERIVE
	 Dim : IfcDimensionCount := HIINDEX(DirectionRatios);
 WHERE
	MagnitudeGreaterZero : SIZEOF(QUERY(Tmp <* DirectionRatios | Tmp <> 0.0)) > 0;
END_ENTITY;`,
	"IFCDIRECTRIXCURVESWEPTAREASOLID": `ENTITY IfcDirectrixCurveSweptAreaSolid
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcFixedReferenceSweptAreaSolid
	,IfcSurfaceCurveSweptAreaSolid))
 SUBTYPE OF (IfcSweptAreaSolid);
	Directrix : IfcCurve;
	StartParam : OPTIONAL IfcCurveMeasureSelect;
	EndParam : OPTIONAL IfcCurveMeasureSelect;
 WHERE
	DirectrixBounded : (EXISTS(StartParam) AND EXISTS(EndParam)) OR 
(SIZEOF(['IFC4X3_DEV_73740fe4.IFCCONIC', 'IFC4X3_DEV_73740fe4.IFCBOUNDEDCURVE'] * TYPEOF(Directrix)) = 1);
END_ENTITY;`,
	"IFCDIRECTRIXDERIVEDREFERENCESWEPTAREASOLID": `ENTITY IfcDirectrixDerivedReferenceSweptAreaSolid
 SUBTYPE OF (IfcFixedReferenceSweptAreaSolid);
END_ENTITY;`,
	"IFCDISCRETEACCESSORY": `ENTITY IfcDiscreteAccessory
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcDiscreteAccessoryTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDiscreteAccessoryTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDiscreteAccessoryTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDISCRETEACCESSORYTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDISCRETEACCESSORYTYPE": `ENTITY IfcDiscreteAccessoryType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcDiscreteAccessoryTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDiscreteAccessoryTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDiscreteAccessoryTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDISTRIBUTIONBOARD": `ENTITY IfcDistributionBoard
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcDistributionBoardTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDistributionBoardTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDistributionBoardTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDISTRIBUTIONBOARDTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDISTRIBUTIONBOARDTYPE": `ENTITY IfcDistributionBoardType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcDistributionBoardTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDistributionBoardTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDistributionBoardTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDISTRIBUTIONCHAMBERELEMENT": `ENTITY IfcDistributionChamberElement
 SUBTYPE OF (IfcDistributionFlowElement);
	PredefinedType : OPTIONAL IfcDistributionChamberElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDistributionChamberElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDistributionChamberElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDISTRIBUTIONCHAMBERELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDISTRIBUTIONCHAMBERELEMENTTYPE": `ENTITY IfcDistributionChamberElementType
 SUBTYPE OF (IfcDistributionFlowElementType);
	PredefinedType : IfcDistributionChamberElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDistributionChamberElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDistributionChamberElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDISTRIBUTIONCIRCUIT": `ENTITY IfcDistributionCircuit
 SUBTYPE OF (IfcDistributionSystem);
END_ENTITY;`,
	"IFCDISTRIBUTIONCONTROLELEMENT": `ENTITY IfcDistributionControlElement
 SUPERTYPE OF (ONEOF
	(IfcActuator
	,IfcAlarm
	,IfcController
	,IfcFlowInstrument
	,IfcProtectiveDeviceTrippingUnit
	,IfcSensor
	,IfcUnitaryControlElement))
 SUBTYPE OF (IfcDistributionElement);
 INVERSE
	AssignedToFlowElement : SET [0:1] OF IfcRelFlowControlElements FOR RelatedControlElements;
END_ENTITY;`,
	"IFCDISTRIBUTIONCONTROLELEMENTTYPE": `ENTITY IfcDistributionControlElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcActuatorType
	,IfcAlarmType
	,IfcControllerType
	,IfcFlowInstrumentType
	,IfcProtectiveDeviceTrippingUnitType
	,IfcSensorType
	,IfcUnitaryControlElementType))
 SUBTYPE OF (IfcDistributionElementType);
END_ENTITY;`,
	"IFCDISTRIBUTIONELEMENT": `ENTITY IfcDistributionElement
 SUPERTYPE OF (ONEOF
	(IfcDistributionControlElement
	,IfcDistributionFlowElement))
 SUBTYPE OF (IfcElement);
 INVERSE
	HasPorts : SET [0:?] OF IfcRelConnectsPortToElement FOR RelatedElement;
END_ENTITY;`,
	"IFCDISTRIBUTIONELEMENTTYPE": `ENTITY IfcDistributionElementType
 SUPERTYPE OF (ONEOF
	(IfcDistributionControlElementType
	,IfcDistributionFlowElementType))
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCDISTRIBUTIONFLOWELEMENT": `ENTITY IfcDistributionFlowElement
 SUPERTYPE OF (ONEOF
	(IfcDistributionChamberElement
	,IfcEnergyConversionDevice
	,IfcFlowController
	,IfcFlowFitting
	,IfcFlowMovingDevice
	,IfcFlowSegment
	,IfcFlowStorageDevice
	,IfcFlowTerminal
	,IfcFlowTreatmentDevice))
 SUBTYPE OF (IfcDistributionElement);
 INVERSE
	HasControlElements : SET [0:1] OF IfcRelFlowControlElements FOR RelatingFlowElement;
END_ENTITY;`,
	"IFCDISTRIBUTIONFLOWELEMENTTYPE": `ENTITY IfcDistributionFlowElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDistributionChamberElementType
	,IfcEnergyConversionDeviceType
	,IfcFlowControllerType
	,IfcFlowFittingType
	,IfcFlowMovingDeviceType
	,IfcFlowSegmentType
	,IfcFlowStorageDeviceType
	,IfcFlowTerminalType
	,IfcFlowTreatmentDeviceType))
 SUBTYPE OF (IfcDistributionElementType);
END_ENTITY;`,
	"IFCDISTRIBUTIONPORT": `ENTITY IfcDistributionPort
 SUBTYPE OF (IfcPort);
	FlowDirection : OPTIONAL IfcFlowDirectionEnum;
	PredefinedType : OPTIONAL IfcDistributionPortTypeEnum;
	SystemType : OPTIONAL IfcDistributionSystemEnum;
END_ENTITY;`,
	"IFCDISTRIBUTIONSYSTEM": `ENTITY IfcDistributionSystem
 SUPERTYPE OF (ONEOF
	(IfcDistributionCircuit))
 SUBTYPE OF (IfcSystem);
	LongName : OPTIONAL IfcLabel;
	PredefinedType : OPTIONAL IfcDistributionSystemEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDistributionSystemEnum.USERDEFINED) OR
 ((PredefinedType = IfcDistributionSystemEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCDOCUMENTINFORMATION": `ENTITY IfcDocumentInformation
 SUBTYPE OF (IfcExternalInformation);
	Identification : IfcIdentifier;
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	Location : OPTIONAL IfcURIReference;
	Purpose : OPTIONAL IfcText;
	IntendedUse : OPTIONAL IfcText;
	Scope : OPTIONAL IfcText;
	Revision : OPTIONAL IfcLabel;
	DocumentOwner : OPTIONAL IfcActorSelect;
	Editors : OPTIONAL SET [1:?] OF IfcActorSelect;
	CreationTime : OPTIONAL IfcDateTime;
	LastRevisionTime : OPTIONAL IfcDateTime;
	ElectronicFormat : OPTIONAL IfcIdentifier;
	ValidFrom : OPTIONAL IfcDate;
	ValidUntil : OPTIONAL IfcDate;
	Confidentiality : OPTIONAL IfcDocumentConfidentialityEnum;
	Status : OPTIONAL IfcDocumentStatusEnum;
 INVERSE
	DocumentInfoForObjects : SET [0:?] OF IfcRelAssociatesDocument FOR RelatingDocument;
	HasDocumentReferences : SET [0:?] OF IfcDocumentReference FOR ReferencedDocument;
	IsPointedTo : SET [0:?] OF IfcDocumentInformationRelationship FOR RelatedDocuments;
	IsPointer : SET [0:1] OF IfcDocumentInformationRelationship FOR RelatingDocument;
END_ENTITY;`,
	"IFCDOCUMENTINFORMATIONRELATIONSHIP": `ENTITY IfcDocumentInformationRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingDocument : IfcDocumentInformation;
	RelatedDocuments : SET [1:?] OF IfcDocumentInformation;
	RelationshipType : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCDOCUMENTREFERENCE": `ENTITY IfcDocumentReference
 SUBTYPE OF (IfcExternalReference);
	Description : OPTIONAL IfcText;
	ReferencedDocument : OPTIONAL IfcDocumentInformation;
 INVERSE
	DocumentRefForObjects : SET [0:?] OF IfcRelAssociatesDocument FOR RelatingDocument;
 WHERE
	WR1 : EXISTS(Name) XOR EXISTS(ReferencedDocument);
END_ENTITY;`,
	"IFCDOOR": `ENTITY IfcDoor
 SUBTYPE OF (IfcBuiltElement);
	OverallHeight : OPTIONAL IfcPositiveLengthMeasure;
	OverallWidth : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcDoorTypeEnum;
	OperationType : OPTIONAL IfcDoorTypeOperationEnum;
	UserDefinedOperationType : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDoorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDoorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDOORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDOORLININGPROPERTIES": `ENTITY IfcDoorLiningProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	LiningDepth : OPTIONAL IfcPositiveLengthMeasure;
	LiningThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	ThresholdDepth : OPTIONAL IfcPositiveLengthMeasure;
	ThresholdThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	TransomThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	TransomOffset : OPTIONAL IfcLengthMeasure;
	LiningOffset : OPTIONAL IfcLengthMeasure;
	ThresholdOffset : OPTIONAL IfcLengthMeasure;
	CasingThickness : OPTIONAL IfcPositiveLengthMeasure;
	CasingDepth : OPTIONAL IfcPositiveLengthMeasure;
	ShapeAspectStyle : OPTIONAL IfcShapeAspect;
	LiningToPanelOffsetX : OPTIONAL IfcLengthMeasure;
	LiningToPanelOffsetY : OPTIONAL IfcLengthMeasure;
 WHERE
	WR31 : NOT(EXISTS(LiningDepth) AND NOT(EXISTS(LiningThickness)));
	WR32 : NOT(EXISTS(ThresholdDepth) AND NOT(EXISTS(ThresholdThickness)));
	WR33 : (EXISTS(TransomOffset) AND EXISTS(TransomThickness)) XOR
(NOT(EXISTS(TransomOffset)) AND NOT(EXISTS(TransomThickness)));
	WR34 : (EXISTS(CasingDepth) AND EXISTS(CasingThickness)) XOR
(NOT(EXISTS(CasingDepth)) AND NOT(EXISTS(CasingThickness)));
	WR35 : (EXISTS(SELF\IfcPropertySetDefinition.DefinesType[1])) 
AND 
('IFC4X3_DEV_73740fe4.IFCDOORTYPE' IN TYPEOF(SELF\IfcPropertySetDefinition.DefinesType[1]));
END_ENTITY;`,
	"IFCDOORPANELPROPERTIES": `ENTITY IfcDoorPanelProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	PanelDepth : OPTIONAL IfcPositiveLengthMeasure;
	PanelOperation : IfcDoorPanelOperationEnum;
	PanelWidth : OPTIONAL IfcNormalisedRatioMeasure;
	PanelPosition : IfcDoorPanelPositionEnum;
	ShapeAspectStyle : OPTIONAL IfcShapeAspect;
 WHERE
	ApplicableToType : (EXISTS(SELF\IfcPropertySetDefinition.DefinesType[1])) 
AND 
('IFC4X3_DEV_73740fe4.IFCDOORTYPE' IN TYPEOF(SELF\IfcPropertySetDefinition.DefinesType[1]));
END_ENTITY;`,
	"IFCDOORTYPE": `ENTITY IfcDoorType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcDoorTypeEnum;
	OperationType : IfcDoorTypeOperationEnum;
	ParameterTakesPrecedence : OPTIONAL IfcBoolean;
	UserDefinedOperationType : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDoorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDoorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDRAUGHTINGPREDEFINEDCOLOUR": `ENTITY IfcDraughtingPreDefinedColour
 SUBTYPE OF (IfcPreDefinedColour);
 WHERE
	PreDefinedColourNames : SELF\IfcPreDefinedItem.Name IN ['black','red','green','blue','yellow',
     'magenta','cyan','white','by layer'];
END_ENTITY;`,
	"IFCDRAUGHTINGPREDEFINEDCURVEFONT": `ENTITY IfcDraughtingPreDefinedCurveFont
 SUBTYPE OF (IfcPreDefinedCurveFont);
 WHERE
	PreDefinedCurveFontNames : SELF\IfcPredefinedItem.Name IN
       ['continuous',
        'chain',
        'chain double dash',
        'dashed',
        'dotted',
        'by layer'];
END_ENTITY;`,
	"IFCDUCTFITTING": `ENTITY IfcDuctFitting
 SUBTYPE OF (IfcFlowFitting);
	PredefinedType : OPTIONAL IfcDuctFittingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDuctFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDUCTFITTINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDUCTFITTINGTYPE": `ENTITY IfcDuctFittingType
 SUBTYPE OF (IfcFlowFittingType);
	PredefinedType : IfcDuctFittingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDuctFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDUCTSEGMENT": `ENTITY IfcDuctSegment
 SUBTYPE OF (IfcFlowSegment);
	PredefinedType : OPTIONAL IfcDuctSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDuctSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDUCTSEGMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDUCTSEGMENTTYPE": `ENTITY IfcDuctSegmentType
 SUBTYPE OF (IfcFlowSegmentType);
	PredefinedType : IfcDuctSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDuctSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCDUCTSILENCER": `ENTITY IfcDuctSilencer
 SUBTYPE OF (IfcFlowTreatmentDevice);
	PredefinedType : OPTIONAL IfcDuctSilencerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcDuctSilencerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctSilencerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCDUCTSILENCERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCDUCTSILENCERTYPE": `ENTITY IfcDuctSilencerType
 SUBTYPE OF (IfcFlowTreatmentDeviceType);
	PredefinedType : IfcDuctSilencerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcDuctSilencerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcDuctSilencerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCEARTHWORKSCUT": `ENTITY IfcEarthworksCut
 SUBTYPE OF (IfcFeatureElementSubtraction);
	PredefinedType : OPTIONAL IfcEarthworksCutTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcEarthworksCutTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEarthworksCutTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCEARTHWORKSELEMENT": `ENTITY IfcEarthworksElement
 SUPERTYPE OF (ONEOF
	(IfcEarthworksFill
	,IfcReinforcedSoil))
 SUBTYPE OF (IfcBuiltElement);
END_ENTITY;`,
	"IFCEARTHWORKSFILL": `ENTITY IfcEarthworksFill
 SUBTYPE OF (IfcEarthworksElement);
	PredefinedType : OPTIONAL IfcEarthworksFillTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcEarthworksFillTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEarthworksFillTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCEDGE": `ENTITY IfcEdge
 SUPERTYPE OF (ONEOF
	(IfcEdgeCurve
	,IfcOrientedEdge
	,IfcSubedge))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
	EdgeStart : IfcVertex;
	EdgeEnd : IfcVertex;
END_ENTITY;`,
	"IFCEDGECURVE": `ENTITY IfcEdgeCurve
 SUBTYPE OF (IfcEdge);
	EdgeGeometry : IfcCurve;
	SameSense : IfcBoolean;
END_ENTITY;`,
	"IFCEDGELOOP": `ENTITY IfcEdgeLoop
 SUBTYPE OF (IfcLoop);
	EdgeList : LIST [1:?] OF IfcOrientedEdge;
 DERIVE
	 Ne : IfcInteger := SIZEOF(EdgeList);
 WHERE
	IsClosed : (EdgeList[1].EdgeStart) :=: (EdgeList[Ne].EdgeEnd);
	IsContinuous : IfcLoopHeadToTail(SELF);
END_ENTITY;`,
	"IFCELECTRICAPPLIANCE": `ENTITY IfcElectricAppliance
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcElectricApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICAPPLIANCETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICAPPLIANCETYPE": `ENTITY IfcElectricApplianceType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcElectricApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICDISTRIBUTIONBOARD": `ENTITY IfcElectricDistributionBoard
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcElectricDistributionBoardTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricDistributionBoardTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricDistributionBoardTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICDISTRIBUTIONBOARDTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICDISTRIBUTIONBOARDTYPE": `ENTITY IfcElectricDistributionBoardType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcElectricDistributionBoardTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricDistributionBoardTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricDistributionBoardTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICFLOWSTORAGEDEVICE": `ENTITY IfcElectricFlowStorageDevice
 SUBTYPE OF (IfcFlowStorageDevice);
	PredefinedType : OPTIONAL IfcElectricFlowStorageDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricFlowStorageDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricFlowStorageDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICFLOWSTORAGEDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICFLOWSTORAGEDEVICETYPE": `ENTITY IfcElectricFlowStorageDeviceType
 SUBTYPE OF (IfcFlowStorageDeviceType);
	PredefinedType : IfcElectricFlowStorageDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricFlowStorageDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricFlowStorageDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICFLOWTREATMENTDEVICE": `ENTITY IfcElectricFlowTreatmentDevice
 SUBTYPE OF (IfcFlowTreatmentDevice);
	PredefinedType : OPTIONAL IfcElectricFlowTreatmentDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricFlowTreatmentDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricFlowTreatmentDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICFLOWTREATMENTDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICFLOWTREATMENTDEVICETYPE": `ENTITY IfcElectricFlowTreatmentDeviceType
 SUBTYPE OF (IfcFlowTreatmentDeviceType);
	PredefinedType : IfcElectricFlowTreatmentDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricFlowTreatmentDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricFlowTreatmentDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICGENERATOR": `ENTITY IfcElectricGenerator
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcElectricGeneratorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricGeneratorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricGeneratorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICGENERATORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICGENERATORTYPE": `ENTITY IfcElectricGeneratorType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcElectricGeneratorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricGeneratorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricGeneratorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICMOTOR": `ENTITY IfcElectricMotor
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcElectricMotorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricMotorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricMotorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICMOTORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICMOTORTYPE": `ENTITY IfcElectricMotorType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcElectricMotorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricMotorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricMotorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELECTRICTIMECONTROL": `ENTITY IfcElectricTimeControl
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcElectricTimeControlTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElectricTimeControlTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricTimeControlTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELECTRICTIMECONTROLTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELECTRICTIMECONTROLTYPE": `ENTITY IfcElectricTimeControlType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcElectricTimeControlTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElectricTimeControlTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElectricTimeControlTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELEMENT": `ENTITY IfcElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBuiltElement
	,IfcCivilElement
	,IfcDistributionElement
	,IfcElementAssembly
	,IfcElementComponent
	,IfcFeatureElement
	,IfcFurnishingElement
	,IfcGeographicElement
	,IfcGeotechnicalElement
	,IfcTransportationDevice
	,IfcVirtualElement))
 SUBTYPE OF (IfcProduct);
	Tag : OPTIONAL IfcIdentifier;
 INVERSE
	FillsVoids : SET [0:1] OF IfcRelFillsElement FOR RelatedBuildingElement;
	ConnectedTo : SET [0:?] OF IfcRelConnectsElements FOR RelatingElement;
	IsInterferedByElements : SET [0:?] OF IfcRelInterferesElements FOR RelatedElement;
	InterferesElements : SET [0:?] OF IfcRelInterferesElements FOR RelatingElement;
	HasProjections : SET [0:?] OF IfcRelProjectsElement FOR RelatingElement;
	HasOpenings : SET [0:?] OF IfcRelVoidsElement FOR RelatingBuildingElement;
	IsConnectionRealization : SET [0:?] OF IfcRelConnectsWithRealizingElements FOR RealizingElements;
	ProvidesBoundaries : SET [0:?] OF IfcRelSpaceBoundary FOR RelatedBuildingElement;
	ConnectedFrom : SET [0:?] OF IfcRelConnectsElements FOR RelatedElement;
	ContainedInStructure : SET [0:1] OF IfcRelContainedInSpatialStructure FOR RelatedElements;
	HasCoverings : SET [0:?] OF IfcRelCoversBldgElements FOR RelatingBuildingElement;
	HasSurfaceFeatures : SET [0:?] OF IfcRelAdheresToElement FOR RelatingElement;
END_ENTITY;`,
	"IFCELEMENTARYSURFACE": `ENTITY IfcElementarySurface
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCylindricalSurface
	,IfcPlane
	,IfcSphericalSurface
	,IfcToroidalSurface))
 SUBTYPE OF (IfcSurface);
	Position : IfcAxis2Placement3D;
END_ENTITY;`,
	"IFCELEMENTASSEMBLY": `ENTITY IfcElementAssembly
 SUBTYPE OF (IfcElement);
	AssemblyPlace : OPTIONAL IfcAssemblyPlaceEnum;
	PredefinedType : OPTIONAL IfcElementAssemblyTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcElementAssemblyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElementAssemblyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCELEMENTASSEMBLYTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCELEMENTASSEMBLYTYPE": `ENTITY IfcElementAssemblyType
 SUBTYPE OF (IfcElementType);
	PredefinedType : IfcElementAssemblyTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcElementAssemblyTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcElementAssemblyTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCELEMENTCOMPONENT": `ENTITY IfcElementComponent
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBuildingElementPart
	,IfcDiscreteAccessory
	,IfcFastener
	,IfcImpactProtectionDevice
	,IfcMechanicalFastener
	,IfcReinforcingElement
	,IfcSign
	,IfcVibrationDamper
	,IfcVibrationIsolator))
 SUBTYPE OF (IfcElement);
END_ENTITY;`,
	"IFCELEMENTCOMPONENTTYPE": `ENTITY IfcElementComponentType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBuildingElementPartType
	,IfcDiscreteAccessoryType
	,IfcFastenerType
	,IfcImpactProtectionDeviceType
	,IfcMechanicalFastenerType
	,IfcReinforcingElementType
	,IfcSignType
	,IfcVibrationDamperType
	,IfcVibrationIsolatorType))
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCELEMENTQUANTITY": `ENTITY IfcElementQuantity
 SUBTYPE OF (IfcQuantitySet);
	MethodOfMeasurement : OPTIONAL IfcLabel;
	Quantities : SET [1:?] OF IfcPhysicalQuantity;
 WHERE
	UniqueQuantityNames : IfcUniqueQuantityNames(Quantities);
END_ENTITY;`,
	"IFCELEMENTTYPE": `ENTITY IfcElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBuiltElementType
	,IfcCivilElementType
	,IfcDistributionElementType
	,IfcElementAssemblyType
	,IfcElementComponentType
	,IfcFurnishingElementType
	,IfcGeographicElementType
	,IfcTransportationDeviceType))
 SUBTYPE OF (IfcTypeProduct);
	ElementType : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCELLIPSE": `ENTITY IfcEllipse
 SUBTYPE OF (IfcConic);
	SemiAxis1 : IfcPositiveLengthMeasure;
	SemiAxis2 : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCELLIPSEPROFILEDEF": `ENTITY IfcEllipseProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	SemiAxis1 : IfcPositiveLengthMeasure;
	SemiAxis2 : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCENERGYCONVERSIONDEVICE": `ENTITY IfcEnergyConversionDevice
 SUPERTYPE OF (ONEOF
	(IfcAirToAirHeatRecovery
	,IfcBoiler
	,IfcBurner
	,IfcChiller
	,IfcCoil
	,IfcCondenser
	,IfcCooledBeam
	,IfcCoolingTower
	,IfcElectricGenerator
	,IfcElectricMotor
	,IfcEngine
	,IfcEvaporativeCooler
	,IfcEvaporator
	,IfcHeatExchanger
	,IfcHumidifier
	,IfcMotorConnection
	,IfcSolarDevice
	,IfcTransformer
	,IfcTubeBundle
	,IfcUnitaryEquipment))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCENERGYCONVERSIONDEVICETYPE": `ENTITY IfcEnergyConversionDeviceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAirToAirHeatRecoveryType
	,IfcBoilerType
	,IfcBurnerType
	,IfcChillerType
	,IfcCoilType
	,IfcCondenserType
	,IfcCooledBeamType
	,IfcCoolingTowerType
	,IfcElectricGeneratorType
	,IfcElectricMotorType
	,IfcEngineType
	,IfcEvaporativeCoolerType
	,IfcEvaporatorType
	,IfcHeatExchangerType
	,IfcHumidifierType
	,IfcMotorConnectionType
	,IfcSolarDeviceType
	,IfcTransformerType
	,IfcTubeBundleType
	,IfcUnitaryEquipmentType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCENGINE": `ENTITY IfcEngine
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcEngineTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcEngineTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEngineTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCENGINETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCENGINETYPE": `ENTITY IfcEngineType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcEngineTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcEngineTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEngineTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCEVAPORATIVECOOLER": `ENTITY IfcEvaporativeCooler
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcEvaporativeCoolerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcEvaporativeCoolerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEvaporativeCoolerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCEVAPORATIVECOOLERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCEVAPORATIVECOOLERTYPE": `ENTITY IfcEvaporativeCoolerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcEvaporativeCoolerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcEvaporativeCoolerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEvaporativeCoolerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCEVAPORATOR": `ENTITY IfcEvaporator
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcEvaporatorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcEvaporatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEvaporatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCEVAPORATORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCEVAPORATORTYPE": `ENTITY IfcEvaporatorType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcEvaporatorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcEvaporatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcEvaporatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCEVENT": `ENTITY IfcEvent
 SUBTYPE OF (IfcProcess);
	PredefinedType : OPTIONAL IfcEventTypeEnum;
	EventTriggerType : OPTIONAL IfcEventTriggerTypeEnum;
	UserDefinedEventTriggerType : OPTIONAL IfcLabel;
	EventOccurenceTime : OPTIONAL IfcEventTime;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcEventTypeEnum.USERDEFINED) OR ((PredefinedType = IfcEventTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : NOT(EXISTS(EventTriggerType)) OR (EventTriggerType <> IfcEventTriggerTypeEnum.USERDEFINED) OR ((EventTriggerType = IfcEventTriggerTypeEnum.USERDEFINED) AND EXISTS(UserDefinedEventTriggerType));
END_ENTITY;`,
	"IFCEVENTTIME": `ENTITY IfcEventTime
 SUBTYPE OF (IfcSchedulingTime);
	ActualDate : OPTIONAL IfcDateTime;
	EarlyDate : OPTIONAL IfcDateTime;
	LateDate : OPTIONAL IfcDateTime;
	ScheduleDate : OPTIONAL IfcDateTime;
END_ENTITY;`,
	"IFCEVENTTYPE": `ENTITY IfcEventType
 SUBTYPE OF (IfcTypeProcess);
	PredefinedType : IfcEventTypeEnum;
	EventTriggerType : IfcEventTriggerTypeEnum;
	UserDefinedEventTriggerType : OPTIONAL IfcLabel;
 WHERE
	CorrectEventTriggerType : (EventTriggerType <> IfcEventTriggerTypeEnum.USERDEFINED) OR ((EventTriggerType = IfcEventTriggerTypeEnum.USERDEFINED) AND EXISTS(UserDefinedEventTriggerType)) ;
	CorrectPredefinedType : (PredefinedType <> IfcEventTypeEnum.USERDEFINED) OR ((PredefinedType = IfcEventTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeProcess.ProcessType));
END_ENTITY;`,
	"IFCEXTENDEDPROPERTIES": `ENTITY IfcExtendedProperties
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMaterialProperties
	,IfcProfileProperties))
 SUBTYPE OF (IfcPropertyAbstraction);
	Name : OPTIONAL IfcIdentifier;
	Description : OPTIONAL IfcText;
	Properties : SET [1:?] OF IfcProperty;
END_ENTITY;`,
	"IFCEXTERNALINFORMATION": `ENTITY IfcExternalInformation
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcClassification
	,IfcDocumentInformation
	,IfcLibraryInformation));
END_ENTITY;`,
	"IFCEXTERNALLYDEFINEDHATCHSTYLE": `ENTITY IfcExternallyDefinedHatchStyle
 SUBTYPE OF (IfcExternalReference);
END_ENTITY;`,
	"IFCEXTERNALLYDEFINEDSURFACESTYLE": `ENTITY IfcExternallyDefinedSurfaceStyle
 SUBTYPE OF (IfcExternalReference);
END_ENTITY;`,
	"IFCEXTERNALLYDEFINEDTEXTFONT": `ENTITY IfcExternallyDefinedTextFont
 SUBTYPE OF (IfcExternalReference);
END_ENTITY;`,
	"IFCEXTERNALREFERENCE": `ENTITY IfcExternalReference
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcClassificationReference
	,IfcDocumentReference
	,IfcExternallyDefinedHatchStyle
	,IfcExternallyDefinedSurfaceStyle
	,IfcExternallyDefinedTextFont
	,IfcLibraryReference));
	Location : OPTIONAL IfcURIReference;
	Identification : OPTIONAL IfcIdentifier;
	Name : OPTIONAL IfcLabel;
 INVERSE
	ExternalReferenceForResources : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatingReference;
 WHERE
	WR1 : EXISTS(Identification) OR EXISTS(Location) OR EXISTS(Name);
END_ENTITY;`,
	"IFCEXTERNALREFERENCERELATIONSHIP": `ENTITY IfcExternalReferenceRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingReference : IfcExternalReference;
	RelatedResourceObjects : SET [1:?] OF IfcResourceObjectSelect;
END_ENTITY;`,
	"IFCEXTERNALSPATIALELEMENT": `ENTITY IfcExternalSpatialElement
 SUBTYPE OF (IfcExternalSpatialStructureElement);
	PredefinedType : OPTIONAL IfcExternalSpatialElementTypeEnum;
 INVERSE
	BoundedBy : SET [0:?] OF IfcRelSpaceBoundary FOR RelatingSpace;
END_ENTITY;`,
	"IFCEXTERNALSPATIALSTRUCTUREELEMENT": `ENTITY IfcExternalSpatialStructureElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcExternalSpatialElement))
 SUBTYPE OF (IfcSpatialElement);
END_ENTITY;`,
	"IFCEXTRUDEDAREASOLID": `ENTITY IfcExtrudedAreaSolid
 SUPERTYPE OF (ONEOF
	(IfcExtrudedAreaSolidTapered))
 SUBTYPE OF (IfcSweptAreaSolid);
	ExtrudedDirection : IfcDirection;
	Depth : IfcPositiveLengthMeasure;
 WHERE
	ValidExtrusionDirection : IfcDotProduct(IfcRepresentationItem() || IfcGeometricRepresentationItem() || IfcDirection([0.0,0.0,1.0]), SELF.ExtrudedDirection) <> 0.0;
END_ENTITY;`,
	"IFCEXTRUDEDAREASOLIDTAPERED": `ENTITY IfcExtrudedAreaSolidTapered
 SUBTYPE OF (IfcExtrudedAreaSolid);
	EndSweptArea : IfcProfileDef;
 WHERE
	CorrectProfileAssignment : IfcTaperedSweptAreaProfiles(SELF\IfcSweptAreaSolid.SweptArea, SELF.EndSweptArea);
END_ENTITY;`,
	"IFCFACE": `ENTITY IfcFace
 SUPERTYPE OF (ONEOF
	(IfcFaceSurface))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
	Bounds : SET [1:?] OF IfcFaceBound;
 INVERSE
	HasTextureMaps : SET [0:?] OF IfcTextureMap FOR MappedTo;
 WHERE
	HasOuterBound : SIZEOF(QUERY(temp <* Bounds | 'IFC4X3_DEV_73740fe4.IFCFACEOUTERBOUND' IN TYPEOF(temp))) <= 1;
END_ENTITY;`,
	"IFCFACEBASEDSURFACEMODEL": `ENTITY IfcFaceBasedSurfaceModel
 SUBTYPE OF (IfcGeometricRepresentationItem);
	FbsmFaces : SET [1:?] OF IfcConnectedFaceSet;
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCFACEBOUND": `ENTITY IfcFaceBound
 SUPERTYPE OF (ONEOF
	(IfcFaceOuterBound))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
	Bound : IfcLoop;
	Orientation : IfcBoolean;
END_ENTITY;`,
	"IFCFACEOUTERBOUND": `ENTITY IfcFaceOuterBound
 SUBTYPE OF (IfcFaceBound);
END_ENTITY;`,
	"IFCFACESURFACE": `ENTITY IfcFaceSurface
 SUPERTYPE OF (ONEOF
	(IfcAdvancedFace))
 SUBTYPE OF (IfcFace);
	FaceSurface : IfcSurface;
	SameSense : IfcBoolean;
END_ENTITY;`,
	"IFCFACETEDBREP": `ENTITY IfcFacetedBrep
 SUPERTYPE OF (ONEOF
	(IfcFacetedBrepWithVoids))
 SUBTYPE OF (IfcManifoldSolidBrep);
END_ENTITY;`,
	"IFCFACETEDBREPWITHVOIDS": `ENTITY IfcFacetedBrepWithVoids
 SUBTYPE OF (IfcFacetedBrep);
	Voids : SET [1:?] OF IfcClosedShell;
END_ENTITY;`,
	"IFCFACILITY": `ENTITY IfcFacility
 SUPERTYPE OF (ONEOF
	(IfcBridge
	,IfcBuilding
	,IfcMarineFacility
	,IfcRailway
	,IfcRoad))
 SUBTYPE OF (IfcSpatialStructureElement);
END_ENTITY;`,
	"IFCFACILITYPART": `ENTITY IfcFacilityPart
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBridgePart
	,IfcFacilityPartCommon
	,IfcMarinePart
	,IfcRailwayPart
	,IfcRoadPart))
 SUBTYPE OF (IfcSpatialStructureElement);
	UsageType : IfcFacilityUsageEnum;
END_ENTITY;`,
	"IFCFACILITYPARTCOMMON": `ENTITY IfcFacilityPartCommon
 SUBTYPE OF (IfcFacilityPart);
	PredefinedType : OPTIONAL IfcFacilityPartCommonTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFacilityPartCommonTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFacilityPartCommonTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCFAILURECONNECTIONCONDITION": `ENTITY IfcFailureConnectionCondition
 SUBTYPE OF (IfcStructuralConnectionCondition);
	TensionFailureX : OPTIONAL IfcForceMeasure;
	TensionFailureY : OPTIONAL IfcForceMeasure;
	TensionFailureZ : OPTIONAL IfcForceMeasure;
	CompressionFailureX : OPTIONAL IfcForceMeasure;
	CompressionFailureY : OPTIONAL IfcForceMeasure;
	CompressionFailureZ : OPTIONAL IfcForceMeasure;
END_ENTITY;`,
	"IFCFAN": `ENTITY IfcFan
 SUBTYPE OF (IfcFlowMovingDevice);
	PredefinedType : OPTIONAL IfcFanTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFanTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFanTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFANTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFANTYPE": `ENTITY IfcFanType
 SUBTYPE OF (IfcFlowMovingDeviceType);
	PredefinedType : IfcFanTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFanTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFanTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFASTENER": `ENTITY IfcFastener
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcFastenerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFastenerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFastenerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFASTENERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFASTENERTYPE": `ENTITY IfcFastenerType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcFastenerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFastenerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFastenerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFEATUREELEMENT": `ENTITY IfcFeatureElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcFeatureElementAddition
	,IfcFeatureElementSubtraction
	,IfcSurfaceFeature))
 SUBTYPE OF (IfcElement);
 WHERE
	NotContained : SIZEOF(ContainedInStructure) = 0;
END_ENTITY;`,
	"IFCFEATUREELEMENTADDITION": `ENTITY IfcFeatureElementAddition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcProjectionElement))
 SUBTYPE OF (IfcFeatureElement);
 INVERSE
	ProjectsElements : IfcRelProjectsElement FOR RelatedFeatureElement;
END_ENTITY;`,
	"IFCFEATUREELEMENTSUBTRACTION": `ENTITY IfcFeatureElementSubtraction
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcEarthworksCut
	,IfcOpeningElement
	,IfcVoidingFeature))
 SUBTYPE OF (IfcFeatureElement);
 INVERSE
	VoidsElements : IfcRelVoidsElement FOR RelatedOpeningElement;
 WHERE
	HasNoSubtraction : SIZEOF(SELF\IfcElement.HasOpenings) = 0;
	IsNotFilling : SIZEOF(SELF\IfcElement.FillsVoids) = 0;
END_ENTITY;`,
	"IFCFILLAREASTYLE": `ENTITY IfcFillAreaStyle
 SUBTYPE OF (IfcPresentationStyle);
	FillStyles : SET [1:?] OF IfcFillStyleSelect;
	ModelOrDraughting : OPTIONAL IfcBoolean;
 WHERE
	ConsistentHatchStyleDef : IfcCorrectFillAreaStyle(SELF.FillStyles);
	MaxOneColour : SIZEOF(QUERY(Style <* SELF.FillStyles |
  'IFC4X3_DEV_73740fe4.IFCCOLOUR' IN
   TYPEOF(Style)
  )) <= 1;
	MaxOneExtHatchStyle : SIZEOF(QUERY(Style <* SELF.FillStyles |
  'IFC4X3_DEV_73740fe4.IFCEXTERNALLYDEFINEDHATCHSTYLE' IN
   TYPEOF(Style)
  )) <= 1;
END_ENTITY;`,
	"IFCFILLAREASTYLEHATCHING": `ENTITY IfcFillAreaStyleHatching
 SUBTYPE OF (IfcGeometricRepresentationItem);
	HatchLineAppearance : IfcCurveStyle;
	StartOfNextHatchLine : IfcHatchLineDistanceSelect;
	PointOfReferenceHatchLine : OPTIONAL IfcCartesianPoint;
	PatternStart : OPTIONAL IfcCartesianPoint;
	HatchLineAngle : IfcPlaneAngleMeasure;
 WHERE
	PatternStart2D : NOT(EXISTS(PatternStart)) OR (PatternStart.Dim = 2)
;
	RefHatchLine2D : NOT(EXISTS(PointOfReferenceHatchLine)) OR (PointOfReferenceHatchLine.Dim = 2);
END_ENTITY;`,
	"IFCFILLAREASTYLETILES": `ENTITY IfcFillAreaStyleTiles
 SUBTYPE OF (IfcGeometricRepresentationItem);
	TilingPattern : LIST [2:2] OF IfcVector;
	Tiles : SET [1:?] OF IfcStyledItem;
	TilingScale : IfcPositiveRatioMeasure;
END_ENTITY;`,
	"IFCFILTER": `ENTITY IfcFilter
 SUBTYPE OF (IfcFlowTreatmentDevice);
	PredefinedType : OPTIONAL IfcFilterTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFilterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFilterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFILTERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFILTERTYPE": `ENTITY IfcFilterType
 SUBTYPE OF (IfcFlowTreatmentDeviceType);
	PredefinedType : IfcFilterTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFilterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFilterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFIRESUPPRESSIONTERMINAL": `ENTITY IfcFireSuppressionTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcFireSuppressionTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFireSuppressionTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFireSuppressionTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFIRESUPPRESSIONTERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFIRESUPPRESSIONTERMINALTYPE": `ENTITY IfcFireSuppressionTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcFireSuppressionTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFireSuppressionTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFireSuppressionTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFIXEDREFERENCESWEPTAREASOLID": `ENTITY IfcFixedReferenceSweptAreaSolid
 SUPERTYPE OF (ONEOF
	(IfcDirectrixDerivedReferenceSweptAreaSolid))
 SUBTYPE OF (IfcDirectrixCurveSweptAreaSolid);
	FixedReference : IfcDirection;
END_ENTITY;`,
	"IFCFLOWCONTROLLER": `ENTITY IfcFlowController
 SUPERTYPE OF (ONEOF
	(IfcAirTerminalBox
	,IfcDamper
	,IfcDistributionBoard
	,IfcElectricDistributionBoard
	,IfcElectricTimeControl
	,IfcFlowMeter
	,IfcProtectiveDevice
	,IfcSwitchingDevice
	,IfcValve))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWCONTROLLERTYPE": `ENTITY IfcFlowControllerType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAirTerminalBoxType
	,IfcDamperType
	,IfcDistributionBoardType
	,IfcElectricDistributionBoardType
	,IfcElectricTimeControlType
	,IfcFlowMeterType
	,IfcProtectiveDeviceType
	,IfcSwitchingDeviceType
	,IfcValveType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWFITTING": `ENTITY IfcFlowFitting
 SUPERTYPE OF (ONEOF
	(IfcCableCarrierFitting
	,IfcCableFitting
	,IfcDuctFitting
	,IfcJunctionBox
	,IfcPipeFitting))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWFITTINGTYPE": `ENTITY IfcFlowFittingType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCableCarrierFittingType
	,IfcCableFittingType
	,IfcDuctFittingType
	,IfcJunctionBoxType
	,IfcPipeFittingType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWINSTRUMENT": `ENTITY IfcFlowInstrument
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcFlowInstrumentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFlowInstrumentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFlowInstrumentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFLOWINSTRUMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFLOWINSTRUMENTTYPE": `ENTITY IfcFlowInstrumentType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcFlowInstrumentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFlowInstrumentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFlowInstrumentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFLOWMETER": `ENTITY IfcFlowMeter
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcFlowMeterTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFlowMeterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFlowMeterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFLOWMETERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFLOWMETERTYPE": `ENTITY IfcFlowMeterType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcFlowMeterTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFlowMeterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFlowMeterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFLOWMOVINGDEVICE": `ENTITY IfcFlowMovingDevice
 SUPERTYPE OF (ONEOF
	(IfcCompressor
	,IfcFan
	,IfcPump))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWMOVINGDEVICETYPE": `ENTITY IfcFlowMovingDeviceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCompressorType
	,IfcFanType
	,IfcPumpType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWSEGMENT": `ENTITY IfcFlowSegment
 SUPERTYPE OF (ONEOF
	(IfcCableCarrierSegment
	,IfcCableSegment
	,IfcConveyorSegment
	,IfcDuctSegment
	,IfcPipeSegment))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWSEGMENTTYPE": `ENTITY IfcFlowSegmentType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCableCarrierSegmentType
	,IfcCableSegmentType
	,IfcConveyorSegmentType
	,IfcDuctSegmentType
	,IfcPipeSegmentType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWSTORAGEDEVICE": `ENTITY IfcFlowStorageDevice
 SUPERTYPE OF (ONEOF
	(IfcElectricFlowStorageDevice
	,IfcTank))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWSTORAGEDEVICETYPE": `ENTITY IfcFlowStorageDeviceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcElectricFlowStorageDeviceType
	,IfcTankType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWTERMINAL": `ENTITY IfcFlowTerminal
 SUPERTYPE OF (ONEOF
	(IfcAirTerminal
	,IfcAudioVisualAppliance
	,IfcCommunicationsAppliance
	,IfcElectricAppliance
	,IfcFireSuppressionTerminal
	,IfcLamp
	,IfcLightFixture
	,IfcLiquidTerminal
	,IfcMedicalDevice
	,IfcMobileTelecommunicationsAppliance
	,IfcOutlet
	,IfcSanitaryTerminal
	,IfcSignal
	,IfcSpaceHeater
	,IfcStackTerminal
	,IfcWasteTerminal))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWTERMINALTYPE": `ENTITY IfcFlowTerminalType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAirTerminalType
	,IfcAudioVisualApplianceType
	,IfcCommunicationsApplianceType
	,IfcElectricApplianceType
	,IfcFireSuppressionTerminalType
	,IfcLampType
	,IfcLightFixtureType
	,IfcLiquidTerminalType
	,IfcMedicalDeviceType
	,IfcMobileTelecommunicationsApplianceType
	,IfcOutletType
	,IfcSanitaryTerminalType
	,IfcSignalType
	,IfcSpaceHeaterType
	,IfcStackTerminalType
	,IfcWasteTerminalType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFLOWTREATMENTDEVICE": `ENTITY IfcFlowTreatmentDevice
 SUPERTYPE OF (ONEOF
	(IfcDuctSilencer
	,IfcElectricFlowTreatmentDevice
	,IfcFilter
	,IfcInterceptor))
 SUBTYPE OF (IfcDistributionFlowElement);
END_ENTITY;`,
	"IFCFLOWTREATMENTDEVICETYPE": `ENTITY IfcFlowTreatmentDeviceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDuctSilencerType
	,IfcElectricFlowTreatmentDeviceType
	,IfcFilterType
	,IfcInterceptorType))
 SUBTYPE OF (IfcDistributionFlowElementType);
END_ENTITY;`,
	"IFCFOOTING": `ENTITY IfcFooting
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcFootingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFootingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFootingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFOOTINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFOOTINGTYPE": `ENTITY IfcFootingType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcFootingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcFootingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFootingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCFURNISHINGELEMENT": `ENTITY IfcFurnishingElement
 SUPERTYPE OF (ONEOF
	(IfcFurniture
	,IfcSystemFurnitureElement))
 SUBTYPE OF (IfcElement);
END_ENTITY;`,
	"IFCFURNISHINGELEMENTTYPE": `ENTITY IfcFurnishingElementType
 SUPERTYPE OF (ONEOF
	(IfcFurnitureType
	,IfcSystemFurnitureElementType))
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCFURNITURE": `ENTITY IfcFurniture
 SUBTYPE OF (IfcFurnishingElement);
	PredefinedType : OPTIONAL IfcFurnitureTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFurnitureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFurnitureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCFURNITURETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCFURNITURETYPE": `ENTITY IfcFurnitureType
 SUBTYPE OF (IfcFurnishingElementType);
	AssemblyPlace : IfcAssemblyPlaceEnum;
	PredefinedType : OPTIONAL IfcFurnitureTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcFurnitureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcFurnitureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCGEOGRAPHICCRS": `ENTITY IfcGeographicCRS
 SUBTYPE OF (IfcCoordinateReferenceSystem);
	PrimeMeridian : OPTIONAL IfcIdentifier;
	AngleUnit : OPTIONAL IfcNamedUnit;
	HeightUnit : OPTIONAL IfcNamedUnit;
 WHERE
	AngleUnitIsPlaneAngle : NOT(EXISTS(SELF.AngleUnit)) OR (SELF.AngleUnit.UnitType=IfcUnitEnum.PLANEANGLEUNIT);
	HeightUnitIsLength : NOT(EXISTS(SELF.HeightUnit)) OR (SELF.HeightUnit.UnitType=IfcUnitEnum.LENGTHUNIT);
END_ENTITY;`,
	"IFCGEOGRAPHICELEMENT": `ENTITY IfcGeographicElement
 SUBTYPE OF (IfcElement);
	PredefinedType : OPTIONAL IfcGeographicElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcGeographicElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcGeographicElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCGEOGRAPHICELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCGEOGRAPHICELEMENTTYPE": `ENTITY IfcGeographicElementType
 SUBTYPE OF (IfcElementType);
	PredefinedType : IfcGeographicElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcGeographicElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcGeographicElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCGEOMETRICCURVESET": `ENTITY IfcGeometricCurveSet
 SUBTYPE OF (IfcGeometricSet);
 WHERE
	NoSurfaces : SIZEOF(QUERY(Temp <* SELF\IfcGeometricSet.Elements |
'IFC4X3_DEV_73740fe4.IFCSURFACE' IN TYPEOF(Temp))) = 0;
END_ENTITY;`,
	"IFCGEOMETRICREPRESENTATIONCONTEXT": `ENTITY IfcGeometricRepresentationContext
 SUPERTYPE OF (ONEOF
	(IfcGeometricRepresentationSubContext))
 SUBTYPE OF (IfcRepresentationContext);
	CoordinateSpaceDimension : IfcDimensionCount;
	Precision : OPTIONAL IfcReal;
	WorldCoordinateSystem : IfcAxis2Placement;
	TrueNorth : OPTIONAL IfcDirection;
 INVERSE
	HasSubContexts : SET [0:?] OF IfcGeometricRepresentationSubContext FOR ParentContext;
	HasCoordinateOperation : SET [0:1] OF IfcCoordinateOperation FOR SourceCRS;
 WHERE
	North2D : NOT(EXISTS(TrueNorth)) OR (HIINDEX(TrueNorth.DirectionRatios) = 2);
END_ENTITY;`,
	"IFCGEOMETRICREPRESENTATIONITEM": `ENTITY IfcGeometricRepresentationItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAnnotationFillArea
	,IfcBooleanResult
	,IfcBoundingBox
	,IfcCartesianPointList
	,IfcCartesianTransformationOperator
	,IfcCsgPrimitive3D
	,IfcCurve
	,IfcDirection
	,IfcFaceBasedSurfaceModel
	,IfcFillAreaStyleHatching
	,IfcFillAreaStyleTiles
	,IfcGeometricSet
	,IfcHalfSpaceSolid
	,IfcLightSource
	,IfcPlacement
	,IfcPlanarExtent
	,IfcPoint
	,IfcSectionedSpine
	,IfcSegment
	,IfcShellBasedSurfaceModel
	,IfcSolidModel
	,IfcSurface
	,IfcTessellatedItem
	,IfcTextLiteral
	,IfcVector))
 SUBTYPE OF (IfcRepresentationItem);
END_ENTITY;`,
	"IFCGEOMETRICREPRESENTATIONSUBCONTEXT": `ENTITY IfcGeometricRepresentationSubContext
 SUBTYPE OF (IfcGeometricRepresentationContext);
	ParentContext : IfcGeometricRepresentationContext;
	TargetScale : OPTIONAL IfcPositiveRatioMeasure;
	TargetView : IfcGeometricProjectionEnum;
	UserDefinedTargetView : OPTIONAL IfcLabel;
 DERIVE
	 SELF\IfcGeometricRepresentationContext.WorldCoordinateSystem : IfcAxis2Placement := ParentContext.WorldCoordinateSystem;
	 SELF\IfcGeometricRepresentationContext.CoordinateSpaceDimension : IfcDimensionCount := ParentContext.CoordinateSpaceDimension;
	 SELF\IfcGeometricRepresentationContext.TrueNorth : IfcDirection := NVL(ParentContext.TrueNorth, IfcConvertDirectionInto2D(SELF\IfcGeometricRepresentationContext.WorldCoordinateSystem.P[2]));
	 SELF\IfcGeometricRepresentationContext.Precision : IfcReal := NVL(ParentContext.Precision,1.E-5);
 WHERE
	NoCoordOperation : SIZEOF(SELF\IfcGeometricRepresentationContext.HasCoordinateOperation) = 0;
	ParentNoSub : NOT('IFC4X3_DEV_73740fe4.IFCGEOMETRICREPRESENTATIONSUBCONTEXT' IN TYPEOF(ParentContext));
	UserTargetProvided :  (TargetView <> IfcGeometricProjectionEnum.USERDEFINED) OR 
((TargetView =  IfcGeometricProjectionEnum.USERDEFINED) AND EXISTS(UserDefinedTargetView)) ;
END_ENTITY;`,
	"IFCGEOMETRICSET": `ENTITY IfcGeometricSet
 SUPERTYPE OF (ONEOF
	(IfcGeometricCurveSet))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Elements : SET [1:?] OF IfcGeometricSetSelect;
 DERIVE
	 Dim : IfcDimensionCount := Elements[1].Dim;
 WHERE
	ConsistentDim : SIZEOF(QUERY(Temp <* Elements |
  Temp.Dim <> Elements[1].Dim))
= 0;
END_ENTITY;`,
	"IFCGEOMODEL": `ENTITY IfcGeomodel
 SUBTYPE OF (IfcGeotechnicalAssembly);
END_ENTITY;`,
	"IFCGEOSLICE": `ENTITY IfcGeoslice
 SUBTYPE OF (IfcGeotechnicalAssembly);
END_ENTITY;`,
	"IFCGEOTECHNICALASSEMBLY": `ENTITY IfcGeotechnicalAssembly
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBorehole
	,IfcGeomodel
	,IfcGeoslice))
 SUBTYPE OF (IfcGeotechnicalElement);
END_ENTITY;`,
	"IFCGEOTECHNICALELEMENT": `ENTITY IfcGeotechnicalElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGeotechnicalAssembly
	,IfcGeotechnicalStratum))
 SUBTYPE OF (IfcElement);
END_ENTITY;`,
	"IFCGEOTECHNICALSTRATUM": `ENTITY IfcGeotechnicalStratum
 SUBTYPE OF (IfcGeotechnicalElement);
	PredefinedType : OPTIONAL IfcGeotechnicalStratumTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcGeotechnicalStratumTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcGeotechnicalStratumTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCGRADIENTCURVE": `ENTITY IfcGradientCurve
 SUBTYPE OF (IfcCompositeCurve);
	BaseCurve : IfcBoundedCurve;
	EndPoint : OPTIONAL IfcPlacement;
END_ENTITY;`,
	"IFCGRID": `ENTITY IfcGrid
 SUBTYPE OF (IfcPositioningElement);
	UAxes : LIST [1:?] OF UNIQUE IfcGridAxis;
	VAxes : LIST [1:?] OF UNIQUE IfcGridAxis;
	WAxes : OPTIONAL LIST [1:?] OF UNIQUE IfcGridAxis;
	PredefinedType : OPTIONAL IfcGridTypeEnum;
END_ENTITY;`,
	"IFCGRIDAXIS": `ENTITY IfcGridAxis;
	AxisTag : OPTIONAL IfcLabel;
	AxisCurve : IfcCurve;
	SameSense : IfcBoolean;
 INVERSE
	PartOfW : SET [0:1] OF IfcGrid FOR WAxes;
	PartOfV : SET [0:1] OF IfcGrid FOR VAxes;
	PartOfU : SET [0:1] OF IfcGrid FOR UAxes;
	HasIntersections : SET [0:?] OF IfcVirtualGridIntersection FOR IntersectingAxes;
 WHERE
	WR1 : AxisCurve.Dim = 2;
	WR2 : (SIZEOF(PartOfU) = 1) XOR (SIZEOF(PartOfV) = 1) XOR (SIZEOF(PartOfW) = 1);
END_ENTITY;`,
	"IFCGRIDPLACEMENT": `ENTITY IfcGridPlacement
 SUBTYPE OF (IfcObjectPlacement);
	PlacementLocation : IfcVirtualGridIntersection;
	PlacementRefDirection : OPTIONAL IfcGridPlacementDirectionSelect;
END_ENTITY;`,
	"IFCGROUP": `ENTITY IfcGroup
 SUPERTYPE OF (ONEOF
	(IfcAsset
	,IfcInventory
	,IfcStructuralLoadGroup
	,IfcStructuralResultGroup
	,IfcSystem))
 SUBTYPE OF (IfcObject);
 INVERSE
	IsGroupedBy : SET [0:?] OF IfcRelAssignsToGroup FOR RelatingGroup;
	ReferencedInStructures : SET [0:?] OF IfcRelReferencedInSpatialStructure FOR RelatedElements;
END_ENTITY;`,
	"IFCHALFSPACESOLID": `ENTITY IfcHalfSpaceSolid
 SUPERTYPE OF (ONEOF
	(IfcBoxedHalfSpace
	,IfcPolygonalBoundedHalfSpace))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	BaseSurface : IfcSurface;
	AgreementFlag : IfcBoolean;
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCHEATEXCHANGER": `ENTITY IfcHeatExchanger
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcHeatExchangerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcHeatExchangerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcHeatExchangerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCHEATEXCHANGERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCHEATEXCHANGERTYPE": `ENTITY IfcHeatExchangerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcHeatExchangerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcHeatExchangerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcHeatExchangerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCHUMIDIFIER": `ENTITY IfcHumidifier
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcHumidifierTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcHumidifierTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcHumidifierTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCHUMIDIFIERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCHUMIDIFIERTYPE": `ENTITY IfcHumidifierType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcHumidifierTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcHumidifierTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcHumidifierTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCIMAGETEXTURE": `ENTITY IfcImageTexture
 SUBTYPE OF (IfcSurfaceTexture);
	URLReference : IfcURIReference;
END_ENTITY;`,
	"IFCIMPACTPROTECTIONDEVICE": `ENTITY IfcImpactProtectionDevice
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcImpactProtectionDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcImpactProtectionDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcImpactProtectionDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCIMPACTPROTECTIONDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCIMPACTPROTECTIONDEVICETYPE": `ENTITY IfcImpactProtectionDeviceType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcImpactProtectionDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcImpactProtectionDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcImpactProtectionDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCINDEXEDCOLOURMAP": `ENTITY IfcIndexedColourMap
 SUBTYPE OF (IfcPresentationItem);
	MappedTo : IfcTessellatedFaceSet;
	Opacity : OPTIONAL IfcNormalisedRatioMeasure;
	Colours : IfcColourRgbList;
	ColourIndex : LIST [1:?] OF IfcPositiveInteger;
END_ENTITY;`,
	"IFCINDEXEDPOLYCURVE": `ENTITY IfcIndexedPolyCurve
 SUBTYPE OF (IfcBoundedCurve);
	Points : IfcCartesianPointList;
	Segments : OPTIONAL LIST [1:?] OF IfcSegmentIndexSelect;
	SelfIntersect : OPTIONAL IfcBoolean;
 WHERE
	Consecutive : NOT(EXISTS(Segments)) OR IfcConsecutiveSegments(Segments);
END_ENTITY;`,
	"IFCINDEXEDPOLYGONALFACE": `ENTITY IfcIndexedPolygonalFace
 SUPERTYPE OF (ONEOF
	(IfcIndexedPolygonalFaceWithVoids))
 SUBTYPE OF (IfcTessellatedItem);
	CoordIndex : LIST [3:?] OF IfcPositiveInteger;
 INVERSE
	ToFaceSet : SET [1:?] OF IfcPolygonalFaceSet FOR Faces;
	HasTexCoords : SET [0:1] OF IfcTextureCoordinateIndices FOR TexCoordsOf;
END_ENTITY;`,
	"IFCINDEXEDPOLYGONALFACEWITHVOIDS": `ENTITY IfcIndexedPolygonalFaceWithVoids
 SUBTYPE OF (IfcIndexedPolygonalFace);
	InnerCoordIndices : LIST [1:?] OF LIST [3:?] OF UNIQUE IfcPositiveInteger;
END_ENTITY;`,
	"IFCINDEXEDPOLYGONALTEXTUREMAP": `ENTITY IfcIndexedPolygonalTextureMap
 SUBTYPE OF (IfcIndexedTextureMap);
	TexCoordIndices : SET [1:?] OF IfcTextureCoordinateIndices;
END_ENTITY;`,
	"IFCINDEXEDTEXTUREMAP": `ENTITY IfcIndexedTextureMap
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcIndexedPolygonalTextureMap
	,IfcIndexedTriangleTextureMap))
 SUBTYPE OF (IfcTextureCoordinate);
	MappedTo : IfcTessellatedFaceSet;
	TexCoords : IfcTextureVertexList;
END_ENTITY;`,
	"IFCINDEXEDTRIANGLETEXTUREMAP": `ENTITY IfcIndexedTriangleTextureMap
 SUBTYPE OF (IfcIndexedTextureMap);
	TexCoordIndex : OPTIONAL LIST [1:?] OF LIST [3:3] OF IfcPositiveInteger;
END_ENTITY;`,
	"IFCINTERCEPTOR": `ENTITY IfcInterceptor
 SUBTYPE OF (IfcFlowTreatmentDevice);
	PredefinedType : OPTIONAL IfcInterceptorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcInterceptorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcInterceptorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCINTERCEPTORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCINTERCEPTORTYPE": `ENTITY IfcInterceptorType
 SUBTYPE OF (IfcFlowTreatmentDeviceType);
	PredefinedType : IfcInterceptorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcInterceptorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcInterceptorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCINTERSECTIONCURVE": `ENTITY IfcIntersectionCurve
 SUBTYPE OF (IfcSurfaceCurve);
 WHERE
	DistinctSurfaces : IfcAssociatedSurface(SELF\IfcSurfaceCurve.AssociatedGeometry[1]) <> IfcAssociatedSurface(SELF\IfcSurfaceCurve.AssociatedGeometry[2]);
	TwoPCurves : SIZEOF(SELF\IfcSurfaceCurve.AssociatedGeometry) = 2;
END_ENTITY;`,
	"IFCINVENTORY": `ENTITY IfcInventory
 SUBTYPE OF (IfcGroup);
	PredefinedType : OPTIONAL IfcInventoryTypeEnum;
	Jurisdiction : OPTIONAL IfcActorSelect;
	ResponsiblePersons : OPTIONAL SET [1:?] OF IfcPerson;
	LastUpdateDate : OPTIONAL IfcDate;
	CurrentValue : OPTIONAL IfcCostValue;
	OriginalValue : OPTIONAL IfcCostValue;
END_ENTITY;`,
	"IFCIRREGULARTIMESERIES": `ENTITY IfcIrregularTimeSeries
 SUBTYPE OF (IfcTimeSeries);
	Values : LIST [1:?] OF IfcIrregularTimeSeriesValue;
END_ENTITY;`,
	"IFCIRREGULARTIMESERIESVALUE": `ENTITY IfcIrregularTimeSeriesValue;
	TimeStamp : IfcDateTime;
	ListValues : LIST [1:?] OF IfcValue;
END_ENTITY;`,
	"IFCISHAPEPROFILEDEF": `ENTITY IfcIShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	OverallWidth : IfcPositiveLengthMeasure;
	OverallDepth : IfcPositiveLengthMeasure;
	WebThickness : IfcPositiveLengthMeasure;
	FlangeThickness : IfcPositiveLengthMeasure;
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	FlangeEdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	FlangeSlope : OPTIONAL IfcPlaneAngleMeasure;
 WHERE
	ValidFilletRadius : NOT(EXISTS(FilletRadius)) OR
((FilletRadius <= (OverallWidth - WebThickness)/2.) AND 
 (FilletRadius <= (OverallDepth - (2. * FlangeThickness))/2.));
	ValidFlangeThickness : (2. * FlangeThickness) < OverallDepth;
	ValidWebThickness : WebThickness < OverallWidth;
END_ENTITY;`,
	"IFCJUNCTIONBOX": `ENTITY IfcJunctionBox
 SUBTYPE OF (IfcFlowFitting);
	PredefinedType : OPTIONAL IfcJunctionBoxTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcJunctionBoxTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcJunctionBoxTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCJUNCTIONBOXTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCJUNCTIONBOXTYPE": `ENTITY IfcJunctionBoxType
 SUBTYPE OF (IfcFlowFittingType);
	PredefinedType : IfcJunctionBoxTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcJunctionBoxTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcJunctionBoxTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCKERB": `ENTITY IfcKerb
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcKerbTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcKerbTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcKerbTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCKERBTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCKERBTYPE": `ENTITY IfcKerbType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcKerbTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcKerbTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcKerbTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCLABORRESOURCE": `ENTITY IfcLaborResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcLaborResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
 (PredefinedType <> IfcLaborResourceTypeEnum.USERDEFINED) OR 
 ((PredefinedType = IfcLaborResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCLABORRESOURCETYPE": `ENTITY IfcLaborResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcLaborResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcLaborResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcLaborResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCLAGTIME": `ENTITY IfcLagTime
 SUBTYPE OF (IfcSchedulingTime);
	LagValue : IfcTimeOrRatioSelect;
	DurationType : IfcTaskDurationEnum;
END_ENTITY;`,
	"IFCLAMP": `ENTITY IfcLamp
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcLampTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcLampTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLampTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCLAMPTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCLAMPTYPE": `ENTITY IfcLampType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcLampTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcLampTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLampTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCLIBRARYINFORMATION": `ENTITY IfcLibraryInformation
 SUBTYPE OF (IfcExternalInformation);
	Name : IfcLabel;
	Version : OPTIONAL IfcLabel;
	Publisher : OPTIONAL IfcActorSelect;
	VersionDate : OPTIONAL IfcDateTime;
	Location : OPTIONAL IfcURIReference;
	Description : OPTIONAL IfcText;
 INVERSE
	LibraryInfoForObjects : SET [0:?] OF IfcRelAssociatesLibrary FOR RelatingLibrary;
	HasLibraryReferences : SET [0:?] OF IfcLibraryReference FOR ReferencedLibrary;
END_ENTITY;`,
	"IFCLIBRARYREFERENCE": `ENTITY IfcLibraryReference
 SUBTYPE OF (IfcExternalReference);
	Description : OPTIONAL IfcText;
	Language : OPTIONAL IfcLanguageId;
	ReferencedLibrary : OPTIONAL IfcLibraryInformation;
 INVERSE
	LibraryRefForObjects : SET [0:?] OF IfcRelAssociatesLibrary FOR RelatingLibrary;
END_ENTITY;`,
	"IFCLIGHTDISTRIBUTIONDATA": `ENTITY IfcLightDistributionData;
	MainPlaneAngle : IfcPlaneAngleMeasure;
	SecondaryPlaneAngle : LIST [1:?] OF IfcPlaneAngleMeasure;
	LuminousIntensity : LIST [1:?] OF IfcLuminousIntensityDistributionMeasure;
END_ENTITY;`,
	"IFCLIGHTFIXTURE": `ENTITY IfcLightFixture
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcLightFixtureTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcLightFixtureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLightFixtureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCLIGHTFIXTURETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCLIGHTFIXTURETYPE": `ENTITY IfcLightFixtureType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcLightFixtureTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcLightFixtureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLightFixtureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCLIGHTINTENSITYDISTRIBUTION": `ENTITY IfcLightIntensityDistribution;
	LightDistributionCurve : IfcLightDistributionCurveEnum;
	DistributionData : LIST [1:?] OF IfcLightDistributionData;
END_ENTITY;`,
	"IFCLIGHTSOURCE": `ENTITY IfcLightSource
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcLightSourceAmbient
	,IfcLightSourceDirectional
	,IfcLightSourceGoniometric
	,IfcLightSourcePositional))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Name : OPTIONAL IfcLabel;
	LightColour : IfcColourRgb;
	AmbientIntensity : OPTIONAL IfcNormalisedRatioMeasure;
	Intensity : OPTIONAL IfcNormalisedRatioMeasure;
END_ENTITY;`,
	"IFCLIGHTSOURCEAMBIENT": `ENTITY IfcLightSourceAmbient
 SUBTYPE OF (IfcLightSource);
END_ENTITY;`,
	"IFCLIGHTSOURCEDIRECTIONAL": `ENTITY IfcLightSourceDirectional
 SUBTYPE OF (IfcLightSource);
	Orientation : IfcDirection;
END_ENTITY;`,
	"IFCLIGHTSOURCEGONIOMETRIC": `ENTITY IfcLightSourceGoniometric
 SUBTYPE OF (IfcLightSource);
	Position : IfcAxis2Placement3D;
	ColourAppearance : OPTIONAL IfcColourRgb;
	ColourTemperature : IfcThermodynamicTemperatureMeasure;
	LuminousFlux : IfcLuminousFluxMeasure;
	LightEmissionSource : IfcLightEmissionSourceEnum;
	LightDistributionDataSource : IfcLightDistributionDataSourceSelect;
END_ENTITY;`,
	"IFCLIGHTSOURCEPOSITIONAL": `ENTITY IfcLightSourcePositional
 SUPERTYPE OF (ONEOF
	(IfcLightSourceSpot))
 SUBTYPE OF (IfcLightSource);
	Position : IfcCartesianPoint;
	Radius : IfcPositiveLengthMeasure;
	ConstantAttenuation : IfcReal;
	DistanceAttenuation : IfcReal;
	QuadricAttenuation : IfcReal;
END_ENTITY;`,
	"IFCLIGHTSOURCESPOT": `ENTITY IfcLightSourceSpot
 SUBTYPE OF (IfcLightSourcePositional);
	Orientation : IfcDirection;
	ConcentrationExponent : OPTIONAL IfcReal;
	SpreadAngle : IfcPositivePlaneAngleMeasure;
	BeamWidthAngle : IfcPositivePlaneAngleMeasure;
END_ENTITY;`,
	"IFCLINE": `ENTITY IfcLine
 SUBTYPE OF (IfcCurve);
	Pnt : IfcCartesianPoint;
	Dir : IfcVector;
 WHERE
	SameDim : Dir.Dim = Pnt.Dim;
END_ENTITY;`,
	"IFCLINEARELEMENT": `ENTITY IfcLinearElement
 SUPERTYPE OF (ONEOF
	(IfcAlignmentCant
	,IfcAlignmentHorizontal
	,IfcAlignmentSegment
	,IfcAlignmentVertical))
 SUBTYPE OF (IfcProduct);
END_ENTITY;`,
	"IFCLINEARPLACEMENT": `ENTITY IfcLinearPlacement
 SUBTYPE OF (IfcObjectPlacement);
	RelativePlacement : IfcAxis2PlacementLinear;
	CartesianPosition : OPTIONAL IfcAxis2Placement3D;
END_ENTITY;`,
	"IFCLINEARPOSITIONINGELEMENT": `ENTITY IfcLinearPositioningElement
 SUPERTYPE OF (ONEOF
	(IfcAlignment))
 SUBTYPE OF (IfcPositioningElement);
END_ENTITY;`,
	"IFCLIQUIDTERMINAL": `ENTITY IfcLiquidTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcLiquidTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcLiquidTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLiquidTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCLIQUIDTERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCLIQUIDTERMINALTYPE": `ENTITY IfcLiquidTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcLiquidTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcLiquidTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcLiquidTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCLOCALPLACEMENT": `ENTITY IfcLocalPlacement
 SUBTYPE OF (IfcObjectPlacement);
	RelativePlacement : IfcAxis2Placement;
 WHERE
	WR21 : IfcCorrectLocalPlacement(RelativePlacement, PlacementRelTo);
END_ENTITY;`,
	"IFCLOOP": `ENTITY IfcLoop
 SUPERTYPE OF (ONEOF
	(IfcEdgeLoop
	,IfcPolyLoop
	,IfcVertexLoop))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
END_ENTITY;`,
	"IFCLSHAPEPROFILEDEF": `ENTITY IfcLShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	Depth : IfcPositiveLengthMeasure;
	Width : OPTIONAL IfcPositiveLengthMeasure;
	Thickness : IfcPositiveLengthMeasure;
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	EdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	LegSlope : OPTIONAL IfcPlaneAngleMeasure;
 WHERE
	ValidThickness : (Thickness < Depth) AND (NOT(EXISTS(Width)) OR (Thickness < Width));
END_ENTITY;`,
	"IFCMANIFOLDSOLIDBREP": `ENTITY IfcManifoldSolidBrep
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAdvancedBrep
	,IfcFacetedBrep))
 SUBTYPE OF (IfcSolidModel);
	Outer : IfcClosedShell;
END_ENTITY;`,
	"IFCMAPCONVERSION": `ENTITY IfcMapConversion
 SUPERTYPE OF (ONEOF
	(IfcMapConversionScaled))
 SUBTYPE OF (IfcCoordinateOperation);
	Eastings : IfcLengthMeasure;
	Northings : IfcLengthMeasure;
	OrthogonalHeight : IfcLengthMeasure;
	XAxisAbscissa : OPTIONAL IfcReal;
	XAxisOrdinate : OPTIONAL IfcReal;
	Scale : OPTIONAL IfcReal;
 WHERE
	TargetCRSOnlyProjected : 'IFC4X3_DEV_73740fe4.IFCPROJECTEDCRS' IN TYPEOF(SELF\IfcCoordinateOperation.TargetCRS);
END_ENTITY;`,
	"IFCMAPCONVERSIONSCALED": `ENTITY IfcMapConversionScaled
 SUBTYPE OF (IfcMapConversion);
	FactorX : IfcReal;
	FactorY : IfcReal;
	FactorZ : IfcReal;
END_ENTITY;`,
	"IFCMAPPEDITEM": `ENTITY IfcMappedItem
 SUBTYPE OF (IfcRepresentationItem);
	MappingSource : IfcRepresentationMap;
	MappingTarget : IfcCartesianTransformationOperator;
END_ENTITY;`,
	"IFCMARINEFACILITY": `ENTITY IfcMarineFacility
 SUBTYPE OF (IfcFacility);
	PredefinedType : OPTIONAL IfcMarineFacilityTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMarineFacilityTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMarineFacilityTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCMARINEPART": `ENTITY IfcMarinePart
 SUBTYPE OF (IfcFacilityPart);
	PredefinedType : OPTIONAL IfcMarinePartTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMarinePartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMarinePartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCMATERIAL": `ENTITY IfcMaterial
 SUBTYPE OF (IfcMaterialDefinition);
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	Category : OPTIONAL IfcLabel;
 INVERSE
	HasRepresentation : SET [0:1] OF IfcMaterialDefinitionRepresentation FOR RepresentedMaterial;
	IsRelatedWith : SET [0:?] OF IfcMaterialRelationship FOR RelatedMaterials;
	RelatesTo : SET [0:1] OF IfcMaterialRelationship FOR RelatingMaterial;
END_ENTITY;`,
	"IFCMATERIALCLASSIFICATIONRELATIONSHIP": `ENTITY IfcMaterialClassificationRelationship;
	MaterialClassifications : SET [1:?] OF IfcClassificationSelect;
	ClassifiedMaterial : IfcMaterial;
END_ENTITY;`,
	"IFCMATERIALCONSTITUENT": `ENTITY IfcMaterialConstituent
 SUBTYPE OF (IfcMaterialDefinition);
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	Material : IfcMaterial;
	Fraction : OPTIONAL IfcNormalisedRatioMeasure;
	Category : OPTIONAL IfcLabel;
 INVERSE
	ToMaterialConstituentSet : IfcMaterialConstituentSet FOR MaterialConstituents;
END_ENTITY;`,
	"IFCMATERIALCONSTITUENTSET": `ENTITY IfcMaterialConstituentSet
 SUBTYPE OF (IfcMaterialDefinition);
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	MaterialConstituents : OPTIONAL SET [1:?] OF IfcMaterialConstituent;
END_ENTITY;`,
	"IFCMATERIALDEFINITION": `ENTITY IfcMaterialDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMaterial
	,IfcMaterialConstituent
	,IfcMaterialConstituentSet
	,IfcMaterialLayer
	,IfcMaterialLayerSet
	,IfcMaterialProfile
	,IfcMaterialProfileSet));
 INVERSE
	AssociatedTo : SET [0:?] OF IfcRelAssociatesMaterial FOR RelatingMaterial;
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
	HasProperties : SET [0:?] OF IfcMaterialProperties FOR Material;
END_ENTITY;`,
	"IFCMATERIALDEFINITIONREPRESENTATION": `ENTITY IfcMaterialDefinitionRepresentation
 SUBTYPE OF (IfcProductRepresentation);
	RepresentedMaterial : IfcMaterial;
 WHERE
	OnlyStyledRepresentations : SIZEOF(QUERY(temp <* Representations | 
  (NOT('IFC4X3_DEV_73740fe4.IFCSTYLEDREPRESENTATION' IN TYPEOF(temp)))
)) = 0;
END_ENTITY;`,
	"IFCMATERIALLAYER": `ENTITY IfcMaterialLayer
 SUPERTYPE OF (ONEOF
	(IfcMaterialLayerWithOffsets))
 SUBTYPE OF (IfcMaterialDefinition);
	Material : OPTIONAL IfcMaterial;
	LayerThickness : IfcNonNegativeLengthMeasure;
	IsVentilated : OPTIONAL IfcLogical;
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	Category : OPTIONAL IfcLabel;
	Priority : OPTIONAL IfcInteger;
 INVERSE
	ToMaterialLayerSet : IfcMaterialLayerSet FOR MaterialLayers;
 WHERE
	NormalizedPriority : NOT(EXISTS(Priority)) OR {0 <= Priority <= 100};
END_ENTITY;`,
	"IFCMATERIALLAYERSET": `ENTITY IfcMaterialLayerSet
 SUBTYPE OF (IfcMaterialDefinition);
	MaterialLayers : LIST [1:?] OF IfcMaterialLayer;
	LayerSetName : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
 DERIVE
	 TotalThickness : IfcLengthMeasure := IfcMlsTotalThickness(SELF);
END_ENTITY;`,
	"IFCMATERIALLAYERSETUSAGE": `ENTITY IfcMaterialLayerSetUsage
 SUBTYPE OF (IfcMaterialUsageDefinition);
	ForLayerSet : IfcMaterialLayerSet;
	LayerSetDirection : IfcLayerSetDirectionEnum;
	DirectionSense : IfcDirectionSenseEnum;
	OffsetFromReferenceLine : IfcLengthMeasure;
	ReferenceExtent : OPTIONAL IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCMATERIALLAYERWITHOFFSETS": `ENTITY IfcMaterialLayerWithOffsets
 SUBTYPE OF (IfcMaterialLayer);
	OffsetDirection : IfcLayerSetDirectionEnum;
	OffsetValues : ARRAY [1:2] OF IfcLengthMeasure;
END_ENTITY;`,
	"IFCMATERIALLIST": `ENTITY IfcMaterialList;
	Materials : LIST [1:?] OF IfcMaterial;
END_ENTITY;`,
	"IFCMATERIALPROFILE": `ENTITY IfcMaterialProfile
 SUPERTYPE OF (ONEOF
	(IfcMaterialProfileWithOffsets))
 SUBTYPE OF (IfcMaterialDefinition);
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	Material : OPTIONAL IfcMaterial;
	Profile : IfcProfileDef;
	Priority : OPTIONAL IfcInteger;
	Category : OPTIONAL IfcLabel;
 INVERSE
	ToMaterialProfileSet : IfcMaterialProfileSet FOR MaterialProfiles;
 WHERE
	NormalizedPriority : NOT(EXISTS(Priority)) OR {0 <= Priority <= 100};
END_ENTITY;`,
	"IFCMATERIALPROFILESET": `ENTITY IfcMaterialProfileSet
 SUBTYPE OF (IfcMaterialDefinition);
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	MaterialProfiles : LIST [1:?] OF IfcMaterialProfile;
	CompositeProfile : OPTIONAL IfcCompositeProfileDef;
END_ENTITY;`,
	"IFCMATERIALPROFILESETUSAGE": `ENTITY IfcMaterialProfileSetUsage
 SUPERTYPE OF (ONEOF
	(IfcMaterialProfileSetUsageTapering))
 SUBTYPE OF (IfcMaterialUsageDefinition);
	ForProfileSet : IfcMaterialProfileSet;
	CardinalPoint : OPTIONAL IfcCardinalPointReference;
	ReferenceExtent : OPTIONAL IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCMATERIALPROFILESETUSAGETAPERING": `ENTITY IfcMaterialProfileSetUsageTapering
 SUBTYPE OF (IfcMaterialProfileSetUsage);
	ForProfileEndSet : IfcMaterialProfileSet;
	CardinalEndPoint : OPTIONAL IfcCardinalPointReference;
END_ENTITY;`,
	"IFCMATERIALPROFILEWITHOFFSETS": `ENTITY IfcMaterialProfileWithOffsets
 SUBTYPE OF (IfcMaterialProfile);
	OffsetValues : ARRAY [1:2] OF IfcLengthMeasure;
END_ENTITY;`,
	"IFCMATERIALPROPERTIES": `ENTITY IfcMaterialProperties
 SUBTYPE OF (IfcExtendedProperties);
	Material : IfcMaterialDefinition;
END_ENTITY;`,
	"IFCMATERIALRELATIONSHIP": `ENTITY IfcMaterialRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingMaterial : IfcMaterial;
	RelatedMaterials : SET [1:?] OF IfcMaterial;
	MaterialExpression : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCMATERIALUSAGEDEFINITION": `ENTITY IfcMaterialUsageDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMaterialLayerSetUsage
	,IfcMaterialProfileSetUsage));
 INVERSE
	AssociatedTo : SET [1:?] OF IfcRelAssociatesMaterial FOR RelatingMaterial;
END_ENTITY;`,
	"IFCMEASUREWITHUNIT": `ENTITY IfcMeasureWithUnit;
	ValueComponent : IfcValue;
	UnitComponent : IfcUnit;
END_ENTITY;`,
	"IFCMECHANICALFASTENER": `ENTITY IfcMechanicalFastener
 SUBTYPE OF (IfcElementComponent);
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	NominalLength : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcMechanicalFastenerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMechanicalFastenerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMechanicalFastenerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMECHANICALFASTENERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMECHANICALFASTENERTYPE": `ENTITY IfcMechanicalFastenerType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcMechanicalFastenerTypeEnum;
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	NominalLength : OPTIONAL IfcPositiveLengthMeasure;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMechanicalFastenerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMechanicalFastenerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCMEDICALDEVICE": `ENTITY IfcMedicalDevice
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcMedicalDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMedicalDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMedicalDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMEDICALDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMEDICALDEVICETYPE": `ENTITY IfcMedicalDeviceType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcMedicalDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMedicalDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMedicalDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCMEMBER": `ENTITY IfcMember
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcMemberTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMemberTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMemberTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMEMBERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMEMBERTYPE": `ENTITY IfcMemberType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcMemberTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMemberTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMemberTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCMETRIC": `ENTITY IfcMetric
 SUBTYPE OF (IfcConstraint);
	Benchmark : IfcBenchmarkEnum;
	ValueSource : OPTIONAL IfcLabel;
	DataValue : OPTIONAL IfcMetricValueSelect;
	ReferencePath : OPTIONAL IfcReference;
END_ENTITY;`,
	"IFCMIRROREDPROFILEDEF": `ENTITY IfcMirroredProfileDef
 SUBTYPE OF (IfcDerivedProfileDef);
 DERIVE
	 SELF\IfcDerivedProfileDef.Operator : IfcCartesianTransformationOperator2D := 
IfcRepresentationItem() || IfcGeometricRepresentationItem() ||
IfcCartesianTransformationOperator(
    -- Axis1
    IfcRepresentationItem() || IfcGeometricRepresentationItem() ||
    IfcDirection([-1., 0.]),
    -- Axis2
    IfcRepresentationItem() || IfcGeometricRepresentationItem() ||
    IfcDirection([ 0., 1.]),
    -- LocalOrigin
    IfcRepresentationItem() || IfcGeometricRepresentationItem() ||
    IfcPoint() || IfcCartesianPoint([0., 0.]),
    -- Scale
    1.) ||
IfcCartesianTransformationOperator2D();
END_ENTITY;`,
	"IFCMOBILETELECOMMUNICATIONSAPPLIANCE": `ENTITY IfcMobileTelecommunicationsAppliance
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcMobileTelecommunicationsApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMobileTelecommunicationsApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMobileTelecommunicationsApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMOBILETELECOMMUNICATIONSAPPLIANCETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMOBILETELECOMMUNICATIONSAPPLIANCETYPE": `ENTITY IfcMobileTelecommunicationsApplianceType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcMobileTelecommunicationsApplianceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMobileTelecommunicationsApplianceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMobileTelecommunicationsApplianceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCMONETARYUNIT": `ENTITY IfcMonetaryUnit;
	Currency : IfcLabel;
END_ENTITY;`,
	"IFCMOORINGDEVICE": `ENTITY IfcMooringDevice
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcMooringDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMooringDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMooringDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMOORINGDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMOORINGDEVICETYPE": `ENTITY IfcMooringDeviceType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcMooringDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMooringDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMooringDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCMOTORCONNECTION": `ENTITY IfcMotorConnection
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcMotorConnectionTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcMotorConnectionTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMotorConnectionTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCMOTORCONNECTIONTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCMOTORCONNECTIONTYPE": `ENTITY IfcMotorConnectionType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcMotorConnectionTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcMotorConnectionTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcMotorConnectionTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCNAMEDUNIT": `ENTITY IfcNamedUnit
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcContextDependentUnit
	,IfcConversionBasedUnit
	,IfcSIUnit));
	Dimensions : IfcDimensionalExponents;
	UnitType : IfcUnitEnum;
 WHERE
	WR1 : IfcCorrectDimensions (SELF.UnitType, SELF.Dimensions);
END_ENTITY;`,
	"IFCNAVIGATIONELEMENT": `ENTITY IfcNavigationElement
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcNavigationElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcNavigationElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcNavigationElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCNAVIGATIONELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCNAVIGATIONELEMENTTYPE": `ENTITY IfcNavigationElementType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcNavigationElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcNavigationElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcNavigationElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCOBJECT": `ENTITY IfcObject
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcActor
	,IfcControl
	,IfcGroup
	,IfcProcess
	,IfcProduct
	,IfcResource))
 SUBTYPE OF (IfcObjectDefinition);
	ObjectType : OPTIONAL IfcLabel;
 INVERSE
	IsDeclaredBy : SET [0:1] OF IfcRelDefinesByObject FOR RelatedObjects;
	Declares : SET [0:?] OF IfcRelDefinesByObject FOR RelatingObject;
	IsTypedBy : SET [0:1] OF IfcRelDefinesByType FOR RelatedObjects;
	IsDefinedBy : SET [0:?] OF IfcRelDefinesByProperties FOR RelatedObjects;
 WHERE
	UniquePropertySetNames : ((SIZEOF(IsDefinedBy) = 0) OR IfcUniqueDefinitionNames(IsDefinedBy));
END_ENTITY;`,
	"IFCOBJECTDEFINITION": `ENTITY IfcObjectDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcContext
	,IfcObject
	,IfcTypeObject))
 SUBTYPE OF (IfcRoot);
 INVERSE
	HasAssignments : SET [0:?] OF IfcRelAssigns FOR RelatedObjects;
	Nests : SET [0:1] OF IfcRelNests FOR RelatedObjects;
	IsNestedBy : SET [0:?] OF IfcRelNests FOR RelatingObject;
	HasContext : SET [0:1] OF IfcRelDeclares FOR RelatedDefinitions;
	IsDecomposedBy : SET [0:?] OF IfcRelAggregates FOR RelatingObject;
	Decomposes : SET [0:1] OF IfcRelAggregates FOR RelatedObjects;
	HasAssociations : SET [0:?] OF IfcRelAssociates FOR RelatedObjects;
END_ENTITY;`,
	"IFCOBJECTIVE": `ENTITY IfcObjective
 SUBTYPE OF (IfcConstraint);
	BenchmarkValues : OPTIONAL LIST [1:?] OF IfcConstraint;
	LogicalAggregator : OPTIONAL IfcLogicalOperatorEnum;
	ObjectiveQualifier : IfcObjectiveEnum;
	UserDefinedQualifier : OPTIONAL IfcLabel;
 WHERE
	WR21 : (ObjectiveQualifier <> IfcObjectiveEnum.USERDEFINED) OR
((ObjectiveQualifier = IfcObjectiveEnum.USERDEFINED) AND EXISTS(SELF\IfcObjective.UserDefinedQualifier));
END_ENTITY;`,
	"IFCOBJECTPLACEMENT": `ENTITY IfcObjectPlacement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGridPlacement
	,IfcLinearPlacement
	,IfcLocalPlacement));
	PlacementRelTo : OPTIONAL IfcObjectPlacement;
 INVERSE
	PlacesObject : SET [0:?] OF IfcProduct FOR ObjectPlacement;
	ReferencedByPlacements : SET [0:?] OF IfcObjectPlacement FOR PlacementRelTo;
END_ENTITY;`,
	"IFCOCCUPANT": `ENTITY IfcOccupant
 SUBTYPE OF (IfcActor);
	PredefinedType : OPTIONAL IfcOccupantTypeEnum;
 WHERE
	WR31 : NOT(PredefinedType = IfcOccupantTypeEnum.USERDEFINED) 
OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCOFFSETCURVE": `ENTITY IfcOffsetCurve
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcOffsetCurve2D
	,IfcOffsetCurve3D
	,IfcOffsetCurveByDistances))
 SUBTYPE OF (IfcCurve);
	BasisCurve : IfcCurve;
END_ENTITY;`,
	"IFCOFFSETCURVE2D": `ENTITY IfcOffsetCurve2D
 SUBTYPE OF (IfcOffsetCurve);
	Distance : IfcLengthMeasure;
	SelfIntersect : IfcLogical;
 WHERE
	DimIs2D : BasisCurve.Dim = 2;
END_ENTITY;`,
	"IFCOFFSETCURVE3D": `ENTITY IfcOffsetCurve3D
 SUBTYPE OF (IfcOffsetCurve);
	Distance : IfcLengthMeasure;
	SelfIntersect : IfcLogical;
	RefDirection : IfcDirection;
 WHERE
	DimIs2D : BasisCurve.Dim = 3;
END_ENTITY;`,
	"IFCOFFSETCURVEBYDISTANCES": `ENTITY IfcOffsetCurveByDistances
 SUBTYPE OF (IfcOffsetCurve);
	OffsetValues : LIST [1:?] OF IfcPointByDistanceExpression;
	Tag : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCOPENCROSSPROFILEDEF": `ENTITY IfcOpenCrossProfileDef
 SUBTYPE OF (IfcProfileDef);
	HorizontalWidths : IfcBoolean;
	Widths : LIST [1:?] OF IfcNonNegativeLengthMeasure;
	Slopes : LIST [1:?] OF IfcPlaneAngleMeasure;
	Tags : OPTIONAL LIST [2:?] OF IfcLabel;
	OffsetPoint : OPTIONAL IfcCartesianPoint;
 WHERE
	CorrectProfileType : SELF\IfcProfileDef.ProfileType = IfcProfileTypeEnum.CURVE;
	CorrespondingSlopeWidths : SIZEOF(Slopes) = SIZEOF(Widths);
	CorrespondingTags : (NOT EXISTS (Tags)) OR (SIZEOF(Tags) = (SIZEOF(Slopes) + 1));
END_ENTITY;`,
	"IFCOPENINGELEMENT": `ENTITY IfcOpeningElement
 SUBTYPE OF (IfcFeatureElementSubtraction);
	PredefinedType : OPTIONAL IfcOpeningElementTypeEnum;
 INVERSE
	HasFillings : SET [0:?] OF IfcRelFillsElement FOR RelatingOpeningElement;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcOpeningElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcOpeningElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCOPENSHELL": `ENTITY IfcOpenShell
 SUBTYPE OF (IfcConnectedFaceSet);
END_ENTITY;`,
	"IFCORGANIZATION": `ENTITY IfcOrganization;
	Identification : OPTIONAL IfcIdentifier;
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	Roles : OPTIONAL LIST [1:?] OF IfcActorRole;
	Addresses : OPTIONAL LIST [1:?] OF IfcAddress;
 INVERSE
	IsRelatedBy : SET [0:?] OF IfcOrganizationRelationship FOR RelatedOrganizations;
	Relates : SET [0:?] OF IfcOrganizationRelationship FOR RelatingOrganization;
	Engages : SET [0:?] OF IfcPersonAndOrganization FOR TheOrganization;
END_ENTITY;`,
	"IFCORGANIZATIONRELATIONSHIP": `ENTITY IfcOrganizationRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingOrganization : IfcOrganization;
	RelatedOrganizations : SET [1:?] OF IfcOrganization;
END_ENTITY;`,
	"IFCORIENTEDEDGE": `ENTITY IfcOrientedEdge
 SUBTYPE OF (IfcEdge);
	EdgeElement : IfcEdge;
	Orientation : IfcBoolean;
 DERIVE
	 SELF\IfcEdge.EdgeStart : IfcVertex := IfcBooleanChoose 
(Orientation, EdgeElement.EdgeStart, EdgeElement.EdgeEnd);
	 SELF\IfcEdge.EdgeEnd : IfcVertex := IfcBooleanChoose 
(Orientation, EdgeElement.EdgeEnd, EdgeElement.EdgeStart);
 WHERE
	EdgeElementNotOriented : NOT('IFC4X3_DEV_73740fe4.IFCORIENTEDEDGE' IN TYPEOF(EdgeElement));
END_ENTITY;`,
	"IFCOUTERBOUNDARYCURVE": `ENTITY IfcOuterBoundaryCurve
 SUBTYPE OF (IfcBoundaryCurve);
END_ENTITY;`,
	"IFCOUTLET": `ENTITY IfcOutlet
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcOutletTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcOutletTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcOutletTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCOUTLETTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCOUTLETTYPE": `ENTITY IfcOutletType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcOutletTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcOutletTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcOutletTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCOWNERHISTORY": `ENTITY IfcOwnerHistory;
	OwningUser : IfcPersonAndOrganization;
	OwningApplication : IfcApplication;
	State : OPTIONAL IfcStateEnum;
	ChangeAction : OPTIONAL IfcChangeActionEnum;
	LastModifiedDate : OPTIONAL IfcTimeStamp;
	LastModifyingUser : OPTIONAL IfcPersonAndOrganization;
	LastModifyingApplication : OPTIONAL IfcApplication;
	CreationDate : IfcTimeStamp;
 WHERE
	CorrectChangeAction : (EXISTS(LastModifiedDate)) OR
(NOT(EXISTS(LastModifiedDate)) AND NOT(EXISTS(ChangeAction))) OR
(NOT(EXISTS(LastModifiedDate)) AND EXISTS(ChangeAction) AND ((ChangeAction = IfcChangeActionEnum.NOTDEFINED) OR (ChangeAction = IfcChangeActionEnum.NOCHANGE)));
END_ENTITY;`,
	"IFCPARAMETERIZEDPROFILEDEF": `ENTITY IfcParameterizedProfileDef
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAsymmetricIShapeProfileDef
	,IfcCShapeProfileDef
	,IfcCircleProfileDef
	,IfcEllipseProfileDef
	,IfcIShapeProfileDef
	,IfcLShapeProfileDef
	,IfcRectangleProfileDef
	,IfcTShapeProfileDef
	,IfcTrapeziumProfileDef
	,IfcUShapeProfileDef
	,IfcZShapeProfileDef))
 SUBTYPE OF (IfcProfileDef);
	Position : OPTIONAL IfcAxis2Placement2D;
END_ENTITY;`,
	"IFCPATH": `ENTITY IfcPath
 SUBTYPE OF (IfcTopologicalRepresentationItem);
	EdgeList : LIST [1:?] OF UNIQUE IfcOrientedEdge;
 WHERE
	IsContinuous : IfcPathHeadToTail(SELF);
END_ENTITY;`,
	"IFCPAVEMENT": `ENTITY IfcPavement
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcPavementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPavementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPavementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPAVEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPAVEMENTTYPE": `ENTITY IfcPavementType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcPavementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPavementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPavementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPCURVE": `ENTITY IfcPcurve
 SUBTYPE OF (IfcCurve);
	BasisSurface : IfcSurface;
	ReferenceCurve : IfcCurve;
 WHERE
	DimIs2D : ReferenceCurve.Dim = 2;
END_ENTITY;`,
	"IFCPERFORMANCEHISTORY": `ENTITY IfcPerformanceHistory
 SUBTYPE OF (IfcControl);
	LifeCyclePhase : IfcLabel;
	PredefinedType : OPTIONAL IfcPerformanceHistoryTypeEnum;
END_ENTITY;`,
	"IFCPERMEABLECOVERINGPROPERTIES": `ENTITY IfcPermeableCoveringProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	OperationType : IfcPermeableCoveringOperationEnum;
	PanelPosition : IfcWindowPanelPositionEnum;
	FrameDepth : OPTIONAL IfcPositiveLengthMeasure;
	FrameThickness : OPTIONAL IfcPositiveLengthMeasure;
	ShapeAspectStyle : OPTIONAL IfcShapeAspect;
END_ENTITY;`,
	"IFCPERMIT": `ENTITY IfcPermit
 SUBTYPE OF (IfcControl);
	PredefinedType : OPTIONAL IfcPermitTypeEnum;
	Status : OPTIONAL IfcLabel;
	LongDescription : OPTIONAL IfcText;
END_ENTITY;`,
	"IFCPERSON": `ENTITY IfcPerson;
	Identification : OPTIONAL IfcIdentifier;
	FamilyName : OPTIONAL IfcLabel;
	GivenName : OPTIONAL IfcLabel;
	MiddleNames : OPTIONAL LIST [1:?] OF IfcLabel;
	PrefixTitles : OPTIONAL LIST [1:?] OF IfcLabel;
	SuffixTitles : OPTIONAL LIST [1:?] OF IfcLabel;
	Roles : OPTIONAL LIST [1:?] OF IfcActorRole;
	Addresses : OPTIONAL LIST [1:?] OF IfcAddress;
 INVERSE
	EngagedIn : SET [0:?] OF IfcPersonAndOrganization FOR ThePerson;
 WHERE
	IdentifiablePerson : EXISTS(Identification) OR EXISTS(FamilyName) OR EXISTS(GivenName);
	ValidSetOfNames : NOT EXISTS(MiddleNames) OR EXISTS(FamilyName) OR EXISTS(GivenName);
END_ENTITY;`,
	"IFCPERSONANDORGANIZATION": `ENTITY IfcPersonAndOrganization;
	ThePerson : IfcPerson;
	TheOrganization : IfcOrganization;
	Roles : OPTIONAL LIST [1:?] OF IfcActorRole;
END_ENTITY;`,
	"IFCPHYSICALCOMPLEXQUANTITY": `ENTITY IfcPhysicalComplexQuantity
 SUBTYPE OF (IfcPhysicalQuantity);
	HasQuantities : SET [1:?] OF IfcPhysicalQuantity;
	Discrimination : IfcLabel;
	Quality : OPTIONAL IfcLabel;
	Usage : OPTIONAL IfcLabel;
 WHERE
	NoSelfReference : SIZEOF(QUERY(temp <* HasQuantities | SELF :=: temp)) = 0;
	UniqueQuantityNames : IfcUniqueQuantityNames(HasQuantities);
END_ENTITY;`,
	"IFCPHYSICALQUANTITY": `ENTITY IfcPhysicalQuantity
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPhysicalComplexQuantity
	,IfcPhysicalSimpleQuantity));
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
 INVERSE
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
	PartOfComplex : SET [0:1] OF IfcPhysicalComplexQuantity FOR HasQuantities;
END_ENTITY;`,
	"IFCPHYSICALSIMPLEQUANTITY": `ENTITY IfcPhysicalSimpleQuantity
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcQuantityArea
	,IfcQuantityCount
	,IfcQuantityLength
	,IfcQuantityNumber
	,IfcQuantityTime
	,IfcQuantityVolume
	,IfcQuantityWeight))
 SUBTYPE OF (IfcPhysicalQuantity);
	Unit : OPTIONAL IfcNamedUnit;
END_ENTITY;`,
	"IFCPILE": `ENTITY IfcPile
 SUBTYPE OF (IfcDeepFoundation);
	PredefinedType : OPTIONAL IfcPileTypeEnum;
	ConstructionType : OPTIONAL IfcPileConstructionEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPileTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPileTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPILETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPILETYPE": `ENTITY IfcPileType
 SUBTYPE OF (IfcDeepFoundationType);
	PredefinedType : IfcPileTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPileTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPileTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPIPEFITTING": `ENTITY IfcPipeFitting
 SUBTYPE OF (IfcFlowFitting);
	PredefinedType : OPTIONAL IfcPipeFittingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPipeFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPipeFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPIPEFITTINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPIPEFITTINGTYPE": `ENTITY IfcPipeFittingType
 SUBTYPE OF (IfcFlowFittingType);
	PredefinedType : IfcPipeFittingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPipeFittingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPipeFittingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPIPESEGMENT": `ENTITY IfcPipeSegment
 SUBTYPE OF (IfcFlowSegment);
	PredefinedType : OPTIONAL IfcPipeSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPipeSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPipeSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPIPESEGMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPIPESEGMENTTYPE": `ENTITY IfcPipeSegmentType
 SUBTYPE OF (IfcFlowSegmentType);
	PredefinedType : IfcPipeSegmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPipeSegmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPipeSegmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPIXELTEXTURE": `ENTITY IfcPixelTexture
 SUBTYPE OF (IfcSurfaceTexture);
	Width : IfcInteger;
	Height : IfcInteger;
	ColourComponents : IfcInteger;
	Pixel : LIST [1:?] OF IfcBinary;
 WHERE
	MinPixelInS : Width >= 1;
	MinPixelInT : Height >= 1;
	NumberOfColours : {1 <= ColourComponents <= 4};
	PixelAsByteAndSameLength : SIZEOF(QUERY(temp<* Pixel |
 (BLENGTH(temp) MOD 8 = 0) AND
 (BLENGTH(temp) = BLENGTH(Pixel[1]))
)) = SIZEOF(Pixel);
	SizeOfPixelList : SIZEOF(Pixel) = (Width * Height);
END_ENTITY;`,
	"IFCPLACEMENT": `ENTITY IfcPlacement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAxis1Placement
	,IfcAxis2Placement2D
	,IfcAxis2Placement3D
	,IfcAxis2PlacementLinear))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Location : IfcPoint;
 DERIVE
	 Dim : IfcDimensionCount := Location.Dim;
END_ENTITY;`,
	"IFCPLANARBOX": `ENTITY IfcPlanarBox
 SUBTYPE OF (IfcPlanarExtent);
	Placement : IfcAxis2Placement;
END_ENTITY;`,
	"IFCPLANAREXTENT": `ENTITY IfcPlanarExtent
 SUPERTYPE OF (ONEOF
	(IfcPlanarBox))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	SizeInX : IfcLengthMeasure;
	SizeInY : IfcLengthMeasure;
END_ENTITY;`,
	"IFCPLANE": `ENTITY IfcPlane
 SUBTYPE OF (IfcElementarySurface);
END_ENTITY;`,
	"IFCPLATE": `ENTITY IfcPlate
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcPlateTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPlateTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPlateTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPLATETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPLATETYPE": `ENTITY IfcPlateType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcPlateTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPlateTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPlateTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPOINT": `ENTITY IfcPoint
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCartesianPoint
	,IfcPointByDistanceExpression
	,IfcPointOnCurve
	,IfcPointOnSurface))
 SUBTYPE OF (IfcGeometricRepresentationItem);
 DERIVE
	 Dim : IfcDimensionCount := IfcPointDim(SELF);
END_ENTITY;`,
	"IFCPOINTBYDISTANCEEXPRESSION": `ENTITY IfcPointByDistanceExpression
 SUBTYPE OF (IfcPoint);
	DistanceAlong : IfcCurveMeasureSelect;
	OffsetLateral : OPTIONAL IfcLengthMeasure;
	OffsetVertical : OPTIONAL IfcLengthMeasure;
	OffsetLongitudinal : OPTIONAL IfcLengthMeasure;
	BasisCurve : IfcCurve;
END_ENTITY;`,
	"IFCPOINTONCURVE": `ENTITY IfcPointOnCurve
 SUBTYPE OF (IfcPoint);
	BasisCurve : IfcCurve;
	PointParameter : IfcParameterValue;
END_ENTITY;`,
	"IFCPOINTONSURFACE": `ENTITY IfcPointOnSurface
 SUBTYPE OF (IfcPoint);
	BasisSurface : IfcSurface;
	PointParameterU : IfcParameterValue;
	PointParameterV : IfcParameterValue;
END_ENTITY;`,
	"IFCPOLYGONALBOUNDEDHALFSPACE": `ENTITY IfcPolygonalBoundedHalfSpace
 SUBTYPE OF (IfcHalfSpaceSolid);
	Position : IfcAxis2Placement3D;
	PolygonalBoundary : IfcBoundedCurve;
 WHERE
	BoundaryDim : PolygonalBoundary.Dim = 2;
	BoundaryType : SIZEOF(TYPEOF(PolygonalBoundary) * [
    'IFC4X3_DEV_73740fe4.IFCPOLYLINE',
    'IFC4X3_DEV_73740fe4.IFCCOMPOSITECURVE',
    'IFC4X3_DEV_73740fe4.IFCINDEXEDPOLYCURVE'
  ]) = 1;
END_ENTITY;`,
	"IFCPOLYGONALFACESET": `ENTITY IfcPolygonalFaceSet
 SUBTYPE OF (IfcTessellatedFaceSet);
	Closed : OPTIONAL IfcBoolean;
	Faces : LIST [1:?] OF UNIQUE IfcIndexedPolygonalFace;
	PnIndex : OPTIONAL LIST [1:?] OF IfcPositiveInteger;
END_ENTITY;`,
	"IFCPOLYLINE": `ENTITY IfcPolyline
 SUBTYPE OF (IfcBoundedCurve);
	Points : LIST [2:?] OF IfcCartesianPoint;
 WHERE
	SameDim : SIZEOF(QUERY(Temp <* Points | Temp.Dim <> Points[1].Dim)) = 0;
END_ENTITY;`,
	"IFCPOLYLOOP": `ENTITY IfcPolyLoop
 SUBTYPE OF (IfcLoop);
	Polygon : LIST [3:?] OF UNIQUE IfcCartesianPoint;
 WHERE
	AllPointsSameDim : SIZEOF(QUERY(Temp <* Polygon | Temp.Dim <> Polygon[1].Dim)) = 0;
END_ENTITY;`,
	"IFCPOLYNOMIALCURVE": `ENTITY IfcPolynomialCurve
 SUBTYPE OF (IfcCurve);
	Position : IfcPlacement;
	CoefficientsX : OPTIONAL LIST [2:?] OF IfcReal;
	CoefficientsY : OPTIONAL LIST [2:?] OF IfcReal;
	CoefficientsZ : OPTIONAL LIST [2:?] OF IfcReal;
 WHERE
	CorrectPositionDim : ((Position.Dim=2) AND (NOT EXISTS(CoefficientsZ))) OR (Position.Dim=3);
	ValidCoefficients : (EXISTS(CoefficientsX) AND EXISTS(CoefficientsY)) OR (EXISTS(CoefficientsX) AND EXISTS(CoefficientsZ)) OR (EXISTS(CoefficientsY) AND EXISTS(CoefficientsZ)) OR (EXISTS(CoefficientsX) AND EXISTS(CoefficientsY) AND EXISTS(CoefficientsZ));
END_ENTITY;`,
	"IFCPORT": `ENTITY IfcPort
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDistributionPort))
 SUBTYPE OF (IfcProduct);
 INVERSE
	ContainedIn : SET [0:1] OF IfcRelConnectsPortToElement FOR RelatingPort;
	ConnectedFrom : SET [0:1] OF IfcRelConnectsPorts FOR RelatedPort;
	ConnectedTo : SET [0:1] OF IfcRelConnectsPorts FOR RelatingPort;
END_ENTITY;`,
	"IFCPOSITIONINGELEMENT": `ENTITY IfcPositioningElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGrid
	,IfcLinearPositioningElement
	,IfcReferent))
 SUBTYPE OF (IfcProduct);
 INVERSE
	ContainedInStructure : SET [0:1] OF IfcRelContainedInSpatialStructure FOR RelatedElements;
	Positions : SET [0:?] OF IfcRelPositions FOR RelatingPositioningElement;
 WHERE
	HasPlacement : EXISTS(SELF\IfcProduct.ObjectPlacement);
END_ENTITY;`,
	"IFCPOSTALADDRESS": `ENTITY IfcPostalAddress
 SUBTYPE OF (IfcAddress);
	InternalLocation : OPTIONAL IfcLabel;
	AddressLines : OPTIONAL LIST [1:?] OF IfcLabel;
	PostalBox : OPTIONAL IfcLabel;
	Town : OPTIONAL IfcLabel;
	Region : OPTIONAL IfcLabel;
	PostalCode : OPTIONAL IfcLabel;
	Country : OPTIONAL IfcLabel;
 WHERE
	WR1 : EXISTS (InternalLocation) OR 
EXISTS (AddressLines) OR
EXISTS (PostalBox) OR
EXISTS (PostalCode) OR
EXISTS (Town) OR 
EXISTS (Region) OR  
EXISTS (Country);
END_ENTITY;`,
	"IFCPREDEFINEDCOLOUR": `ENTITY IfcPreDefinedColour
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDraughtingPreDefinedColour))
 SUBTYPE OF (IfcPreDefinedItem);
END_ENTITY;`,
	"IFCPREDEFINEDCURVEFONT": `ENTITY IfcPreDefinedCurveFont
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDraughtingPreDefinedCurveFont))
 SUBTYPE OF (IfcPreDefinedItem);
END_ENTITY;`,
	"IFCPREDEFINEDITEM": `ENTITY IfcPreDefinedItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPreDefinedColour
	,IfcPreDefinedCurveFont
	,IfcPreDefinedTextFont))
 SUBTYPE OF (IfcPresentationItem);
	Name : IfcLabel;
END_ENTITY;`,
	"IFCPREDEFINEDPROPERTIES": `ENTITY IfcPreDefinedProperties
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcReinforcementBarProperties
	,IfcSectionProperties
	,IfcSectionReinforcementProperties))
 SUBTYPE OF (IfcPropertyAbstraction);
END_ENTITY;`,
	"IFCPREDEFINEDPROPERTYSET": `ENTITY IfcPreDefinedPropertySet
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDoorLiningProperties
	,IfcDoorPanelProperties
	,IfcPermeableCoveringProperties
	,IfcReinforcementDefinitionProperties
	,IfcWindowLiningProperties
	,IfcWindowPanelProperties))
 SUBTYPE OF (IfcPropertySetDefinition);
END_ENTITY;`,
	"IFCPREDEFINEDTEXTFONT": `ENTITY IfcPreDefinedTextFont
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcTextStyleFontModel))
 SUBTYPE OF (IfcPreDefinedItem);
END_ENTITY;`,
	"IFCPRESENTATIONITEM": `ENTITY IfcPresentationItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcColourRgbList
	,IfcColourSpecification
	,IfcCurveStyleFont
	,IfcCurveStyleFontAndScaling
	,IfcCurveStyleFontPattern
	,IfcIndexedColourMap
	,IfcPreDefinedItem
	,IfcSurfaceStyleLighting
	,IfcSurfaceStyleRefraction
	,IfcSurfaceStyleShading
	,IfcSurfaceStyleWithTextures
	,IfcSurfaceTexture
	,IfcTextStyleForDefinedFont
	,IfcTextStyleTextModel
	,IfcTextureCoordinate
	,IfcTextureVertex
	,IfcTextureVertexList));
END_ENTITY;`,
	"IFCPRESENTATIONLAYERASSIGNMENT": `ENTITY IfcPresentationLayerAssignment
 SUPERTYPE OF (ONEOF
	(IfcPresentationLayerWithStyle));
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	AssignedItems : SET [1:?] OF IfcLayeredItem;
	Identifier : OPTIONAL IfcIdentifier;
 WHERE
	ApplicableItems : SIZEOF(QUERY(temp <* AssignedItems | (
  SIZEOF(TYPEOF(temp) * [
    'IFC4X3_DEV_73740fe4.IFCSHAPEREPRESENTATION',
    'IFC4X3_DEV_73740fe4.IFCGEOMETRICREPRESENTATIONITEM',
    'IFC4X3_DEV_73740fe4.IFCMAPPEDITEM']) = 1)
)) = SIZEOF(AssignedItems);
END_ENTITY;`,
	"IFCPRESENTATIONLAYERWITHSTYLE": `ENTITY IfcPresentationLayerWithStyle
 SUBTYPE OF (IfcPresentationLayerAssignment);
	LayerOn : IfcLogical;
	LayerFrozen : IfcLogical;
	LayerBlocked : IfcLogical;
	LayerStyles : SET [0:?] OF IfcPresentationStyle;
 WHERE
	ApplicableOnlyToItems : SIZEOF(QUERY(temp <* AssignedItems | (SIZEOF(TYPEOF(temp) * ['IFC4X3_DEV_73740fe4.IfcGeometricRepresentationItem','IFC4X3_DEV_73740fe4.IfcMappedItem']) >= 1))) = SIZEOF(AssignedItems);
END_ENTITY;`,
	"IFCPRESENTATIONSTYLE": `ENTITY IfcPresentationStyle
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCurveStyle
	,IfcFillAreaStyle
	,IfcSurfaceStyle
	,IfcTextStyle));
	Name : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCPROCEDURE": `ENTITY IfcProcedure
 SUBTYPE OF (IfcProcess);
	PredefinedType : OPTIONAL IfcProcedureTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcProcedureTypeEnum.USERDEFINED) OR
((PredefinedType = IfcProcedureTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
	HasName : EXISTS(SELF\IfcRoot.Name);
END_ENTITY;`,
	"IFCPROCEDURETYPE": `ENTITY IfcProcedureType
 SUBTYPE OF (IfcTypeProcess);
	PredefinedType : IfcProcedureTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcProcedureTypeEnum.USERDEFINED) OR ((PredefinedType = IfcProcedureTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeProcess.ProcessType));
END_ENTITY;`,
	"IFCPROCESS": `ENTITY IfcProcess
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcEvent
	,IfcProcedure
	,IfcTask))
 SUBTYPE OF (IfcObject);
	Identification : OPTIONAL IfcIdentifier;
	LongDescription : OPTIONAL IfcText;
 INVERSE
	IsPredecessorTo : SET [0:?] OF IfcRelSequence FOR RelatingProcess;
	IsSuccessorFrom : SET [0:?] OF IfcRelSequence FOR RelatedProcess;
	OperatesOn : SET [0:?] OF IfcRelAssignsToProcess FOR RelatingProcess;
END_ENTITY;`,
	"IFCPRODUCT": `ENTITY IfcProduct
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcAnnotation
	,IfcElement
	,IfcLinearElement
	,IfcPort
	,IfcPositioningElement
	,IfcSpatialElement
	,IfcStructuralActivity
	,IfcStructuralItem))
 SUBTYPE OF (IfcObject);
	ObjectPlacement : OPTIONAL IfcObjectPlacement;
	Representation : OPTIONAL IfcProductRepresentation;
 INVERSE
	ReferencedBy : SET [0:?] OF IfcRelAssignsToProduct FOR RelatingProduct;
	PositionedRelativeTo : SET [0:?] OF IfcRelPositions FOR RelatedProducts;
	ReferencedInStructures : SET [0:?] OF IfcRelReferencedInSpatialStructure FOR RelatedElements;
 WHERE
	PlacementForShapeRepresentation : (EXISTS(Representation) AND EXISTS(ObjectPlacement))
            OR (EXISTS(Representation) AND 
      (SIZEOF(QUERY(temp <* Representation.Representations | 'IFC4X3_DEV_73740fe4.IFCSHAPEREPRESENTATION' IN TYPEOF(temp))) = 0))
            OR (NOT(EXISTS(Representation)));
END_ENTITY;`,
	"IFCPRODUCTDEFINITIONSHAPE": `ENTITY IfcProductDefinitionShape
 SUBTYPE OF (IfcProductRepresentation);
 INVERSE
	ShapeOfProduct : SET [1:?] OF IfcProduct FOR Representation;
	HasShapeAspects : SET [0:?] OF IfcShapeAspect FOR PartOfProductDefinitionShape;
 WHERE
	OnlyShapeModel : SIZEOF(QUERY(temp <* Representations | 
  (NOT('IFC4X3_DEV_73740fe4.IFCSHAPEMODEL' IN TYPEOF(temp)))
)) = 0;
END_ENTITY;`,
	"IFCPRODUCTREPRESENTATION": `ENTITY IfcProductRepresentation
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcMaterialDefinitionRepresentation
	,IfcProductDefinitionShape));
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	Representations : LIST [1:?] OF IfcRepresentation;
END_ENTITY;`,
	"IFCPROFILEDEF": `ENTITY IfcProfileDef
 SUPERTYPE OF (ONEOF
	(IfcArbitraryClosedProfileDef
	,IfcArbitraryOpenProfileDef
	,IfcCompositeProfileDef
	,IfcDerivedProfileDef
	,IfcOpenCrossProfileDef
	,IfcParameterizedProfileDef));
	ProfileType : IfcProfileTypeEnum;
	ProfileName : OPTIONAL IfcLabel;
 INVERSE
	HasExternalReference : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
	HasProperties : SET [0:?] OF IfcProfileProperties FOR ProfileDefinition;
END_ENTITY;`,
	"IFCPROFILEPROPERTIES": `ENTITY IfcProfileProperties
 SUBTYPE OF (IfcExtendedProperties);
	ProfileDefinition : IfcProfileDef;
END_ENTITY;`,
	"IFCPROJECT": `ENTITY IfcProject
 SUBTYPE OF (IfcContext);
 WHERE
	CorrectContext : NOT(EXISTS(SELF\IfcContext.RepresentationContexts)) OR
(SIZEOF(QUERY(Temp <* SELF\IfcContext.RepresentationContexts |
 'IFC4X3_DEV_73740fe4.IFCGEOMETRICREPRESENTATIONSUBCONTEXT' IN TYPEOF(Temp)
  )) = 0);
	HasName : EXISTS(SELF\IfcRoot.Name);
	NoDecomposition : SIZEOF(SELF\IfcObjectDefinition.Decomposes) = 0;
END_ENTITY;`,
	"IFCPROJECTEDCRS": `ENTITY IfcProjectedCRS
 SUBTYPE OF (IfcCoordinateReferenceSystem);
	VerticalDatum : OPTIONAL IfcIdentifier;
	MapProjection : OPTIONAL IfcIdentifier;
	MapZone : OPTIONAL IfcIdentifier;
	MapUnit : OPTIONAL IfcNamedUnit;
 WHERE
	MapUnitIsLength : NOT(EXISTS(SELF.MapUnit)) OR (SELF.MapUnit.UnitType = IfcUnitEnum.LENGTHUNIT);
END_ENTITY;`,
	"IFCPROJECTIONELEMENT": `ENTITY IfcProjectionElement
 SUBTYPE OF (IfcFeatureElementAddition);
	PredefinedType : OPTIONAL IfcProjectionElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcProjectionElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcProjectionElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCPROJECTLIBRARY": `ENTITY IfcProjectLibrary
 SUBTYPE OF (IfcContext);
END_ENTITY;`,
	"IFCPROJECTORDER": `ENTITY IfcProjectOrder
 SUBTYPE OF (IfcControl);
	PredefinedType : OPTIONAL IfcProjectOrderTypeEnum;
	Status : OPTIONAL IfcLabel;
	LongDescription : OPTIONAL IfcText;
END_ENTITY;`,
	"IFCPROPERTY": `ENTITY IfcProperty
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcComplexProperty
	,IfcSimpleProperty))
 SUBTYPE OF (IfcPropertyAbstraction);
	Name : IfcIdentifier;
	Specification : OPTIONAL IfcText;
 INVERSE
	PartOfPset : SET [0:?] OF IfcPropertySet FOR HasProperties;
	PropertyForDependance : SET [0:?] OF IfcPropertyDependencyRelationship FOR DependingProperty;
	PropertyDependsOn : SET [0:?] OF IfcPropertyDependencyRelationship FOR DependantProperty;
	PartOfComplex : SET [0:?] OF IfcComplexProperty FOR HasProperties;
	HasConstraints : SET [0:?] OF IfcResourceConstraintRelationship FOR RelatedResourceObjects;
	HasApprovals : SET [0:?] OF IfcResourceApprovalRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCPROPERTYABSTRACTION": `ENTITY IfcPropertyAbstraction
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcExtendedProperties
	,IfcPreDefinedProperties
	,IfcProperty
	,IfcPropertyEnumeration));
 INVERSE
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCPROPERTYBOUNDEDVALUE": `ENTITY IfcPropertyBoundedValue
 SUBTYPE OF (IfcSimpleProperty);
	UpperBoundValue : OPTIONAL IfcValue;
	LowerBoundValue : OPTIONAL IfcValue;
	Unit : OPTIONAL IfcUnit;
	SetPointValue : OPTIONAL IfcValue;
 WHERE
	SameUnitLowerSet : NOT(EXISTS(LowerBoundValue)) OR NOT(EXISTS(SetPointValue)) OR
(TYPEOF(LowerBoundValue) = TYPEOF(SetPointValue));
	SameUnitUpperLower : NOT(EXISTS(UpperBoundValue)) OR NOT(EXISTS(LowerBoundValue)) OR
(TYPEOF(UpperBoundValue) = TYPEOF(LowerBoundValue));
	SameUnitUpperSet : NOT(EXISTS(UpperBoundValue)) OR NOT(EXISTS(SetPointValue)) OR
(TYPEOF(UpperBoundValue) = TYPEOF(SetPointValue));
END_ENTITY;`,
	"IFCPROPERTYDEFINITION": `ENTITY IfcPropertyDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPropertySetDefinition
	,IfcPropertyTemplateDefinition))
 SUBTYPE OF (IfcRoot);
 INVERSE
	HasContext : SET [0:1] OF IfcRelDeclares FOR RelatedDefinitions;
	HasAssociations : SET [0:?] OF IfcRelAssociates FOR RelatedObjects;
END_ENTITY;`,
	"IFCPROPERTYDEPENDENCYRELATIONSHIP": `ENTITY IfcPropertyDependencyRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	DependingProperty : IfcProperty;
	DependantProperty : IfcProperty;
	Expression : OPTIONAL IfcText;
 WHERE
	NoSelfReference : DependingProperty :<>: DependantProperty;
END_ENTITY;`,
	"IFCPROPERTYENUMERATEDVALUE": `ENTITY IfcPropertyEnumeratedValue
 SUBTYPE OF (IfcSimpleProperty);
	EnumerationValues : OPTIONAL LIST [1:?] OF IfcValue;
	EnumerationReference : OPTIONAL IfcPropertyEnumeration;
 WHERE
	WR21 : NOT(EXISTS(EnumerationReference)) 
OR  NOT(EXISTS(EnumerationValues)) 
OR  (SIZEOF(QUERY(temp <* EnumerationValues |
    temp IN EnumerationReference.EnumerationValues))
    = SIZEOF(EnumerationValues));
END_ENTITY;`,
	"IFCPROPERTYENUMERATION": `ENTITY IfcPropertyEnumeration
 SUBTYPE OF (IfcPropertyAbstraction);
	Name : IfcLabel;
	EnumerationValues : LIST [1:?] OF UNIQUE IfcValue;
	Unit : OPTIONAL IfcUnit;
 UNIQUE
	UR1 : Name;
 WHERE
	WR01 : SIZEOF(QUERY(temp <* SELF.EnumerationValues | 
  NOT(TYPEOF(SELF.EnumerationValues[1]) = TYPEOF(temp))
  )) = 0;
END_ENTITY;`,
	"IFCPROPERTYLISTVALUE": `ENTITY IfcPropertyListValue
 SUBTYPE OF (IfcSimpleProperty);
	ListValues : OPTIONAL LIST [1:?] OF IfcValue;
	Unit : OPTIONAL IfcUnit;
 WHERE
	WR31 : SIZEOF(QUERY(temp <* SELF.ListValues | 
  NOT(TYPEOF(SELF.ListValues[1]) = TYPEOF(temp))
  )) = 0;
END_ENTITY;`,
	"IFCPROPERTYREFERENCEVALUE": `ENTITY IfcPropertyReferenceValue
 SUBTYPE OF (IfcSimpleProperty);
	UsageName : OPTIONAL IfcText;
	PropertyReference : OPTIONAL IfcObjectReferenceSelect;
END_ENTITY;`,
	"IFCPROPERTYSET": `ENTITY IfcPropertySet
 SUBTYPE OF (IfcPropertySetDefinition);
	HasProperties : SET [1:?] OF IfcProperty;
 WHERE
	ExistsName : EXISTS(SELF\IfcRoot.Name);
	UniquePropertyNames : IfcUniquePropertyName(HasProperties);
END_ENTITY;`,
	"IFCPROPERTYSETDEFINITION": `ENTITY IfcPropertySetDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPreDefinedPropertySet
	,IfcPropertySet
	,IfcQuantitySet))
 SUBTYPE OF (IfcPropertyDefinition);
 INVERSE
	DefinesType : SET [0:?] OF IfcTypeObject FOR HasPropertySets;
	IsDefinedBy : SET [0:?] OF IfcRelDefinesByTemplate FOR RelatedPropertySets;
	DefinesOccurrence : SET [0:?] OF IfcRelDefinesByProperties FOR RelatingPropertyDefinition;
END_ENTITY;`,
	"IFCPROPERTYSETTEMPLATE": `ENTITY IfcPropertySetTemplate
 SUBTYPE OF (IfcPropertyTemplateDefinition);
	TemplateType : OPTIONAL IfcPropertySetTemplateTypeEnum;
	ApplicableEntity : OPTIONAL IfcIdentifier;
	HasPropertyTemplates : SET [1:?] OF IfcPropertyTemplate;
 INVERSE
	Defines : SET [0:?] OF IfcRelDefinesByTemplate FOR RelatingTemplate;
 WHERE
	ExistsName : EXISTS(SELF\IfcRoot.Name);
	UniquePropertyNames : IfcUniquePropertyTemplateNames(HasPropertyTemplates);
END_ENTITY;`,
	"IFCPROPERTYSINGLEVALUE": `ENTITY IfcPropertySingleValue
 SUBTYPE OF (IfcSimpleProperty);
	NominalValue : OPTIONAL IfcValue;
	Unit : OPTIONAL IfcUnit;
END_ENTITY;`,
	"IFCPROPERTYTABLEVALUE": `ENTITY IfcPropertyTableValue
 SUBTYPE OF (IfcSimpleProperty);
	DefiningValues : OPTIONAL LIST [1:?] OF UNIQUE IfcValue;
	DefinedValues : OPTIONAL LIST [1:?] OF IfcValue;
	Expression : OPTIONAL IfcText;
	DefiningUnit : OPTIONAL IfcUnit;
	DefinedUnit : OPTIONAL IfcUnit;
	CurveInterpolation : OPTIONAL IfcCurveInterpolationEnum;
 WHERE
	WR21 : (NOT(EXISTS(DefiningValues)) AND NOT(EXISTS(DefinedValues)))
OR (SIZEOF(DefiningValues) = SIZEOF(DefinedValues));
	WR22 : NOT(EXISTS(DefiningValues)) OR 
(SIZEOF(QUERY(temp <* SELF.DefiningValues | TYPEOF(temp) <> TYPEOF(SELF.DefiningValues[1])
)) = 0);
	WR23 : NOT(EXISTS(DefinedValues)) OR 
(SIZEOF(QUERY(temp <* SELF.DefinedValues | TYPEOF(temp) <> TYPEOF(SELF.DefinedValues[1])
)) = 0);
END_ENTITY;`,
	"IFCPROPERTYTEMPLATE": `ENTITY IfcPropertyTemplate
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcComplexPropertyTemplate
	,IfcSimplePropertyTemplate))
 SUBTYPE OF (IfcPropertyTemplateDefinition);
 INVERSE
	PartOfComplexTemplate : SET [0:?] OF IfcComplexPropertyTemplate FOR HasPropertyTemplates;
	PartOfPsetTemplate : SET [0:?] OF IfcPropertySetTemplate FOR HasPropertyTemplates;
END_ENTITY;`,
	"IFCPROPERTYTEMPLATEDEFINITION": `ENTITY IfcPropertyTemplateDefinition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPropertySetTemplate
	,IfcPropertyTemplate))
 SUBTYPE OF (IfcPropertyDefinition);
END_ENTITY;`,
	"IFCPROTECTIVEDEVICE": `ENTITY IfcProtectiveDevice
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcProtectiveDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcProtectiveDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcProtectiveDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPROTECTIVEDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPROTECTIVEDEVICETRIPPINGUNIT": `ENTITY IfcProtectiveDeviceTrippingUnit
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcProtectiveDeviceTrippingUnitTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcProtectiveDeviceTrippingUnitTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcProtectiveDeviceTrippingUnitTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPROTECTIVEDEVICETRIPPINGUNITTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPROTECTIVEDEVICETRIPPINGUNITTYPE": `ENTITY IfcProtectiveDeviceTrippingUnitType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcProtectiveDeviceTrippingUnitTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcProtectiveDeviceTrippingUnitTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcProtectiveDeviceTrippingUnitTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPROTECTIVEDEVICETYPE": `ENTITY IfcProtectiveDeviceType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcProtectiveDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcProtectiveDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcProtectiveDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCPUMP": `ENTITY IfcPump
 SUBTYPE OF (IfcFlowMovingDevice);
	PredefinedType : OPTIONAL IfcPumpTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcPumpTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPumpTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCPUMPTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCPUMPTYPE": `ENTITY IfcPumpType
 SUBTYPE OF (IfcFlowMovingDeviceType);
	PredefinedType : IfcPumpTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcPumpTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcPumpTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCQUANTITYAREA": `ENTITY IfcQuantityArea
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	AreaValue : IfcAreaMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : NOT(EXISTS(SELF\IfcPhysicalSimpleQuantity.Unit)) OR
   (SELF\IfcPhysicalSimpleQuantity.Unit.UnitType = IfcUnitEnum.AREAUNIT);
	WR22 : AreaValue >= 0.;
END_ENTITY;`,
	"IFCQUANTITYCOUNT": `ENTITY IfcQuantityCount
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	CountValue : IfcCountMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : CountValue >= 0;
END_ENTITY;`,
	"IFCQUANTITYLENGTH": `ENTITY IfcQuantityLength
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	LengthValue : IfcLengthMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : NOT(EXISTS(SELF\IfcPhysicalSimpleQuantity.Unit)) OR
   (SELF\IfcPhysicalSimpleQuantity.Unit.UnitType = IfcUnitEnum.LENGTHUNIT);
	WR22 : LengthValue >= 0.;
END_ENTITY;`,
	"IFCQUANTITYNUMBER": `ENTITY IfcQuantityNumber
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	NumberValue : IfcNumericMeasure;
	Formula : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCQUANTITYSET": `ENTITY IfcQuantitySet
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcElementQuantity))
 SUBTYPE OF (IfcPropertySetDefinition);
END_ENTITY;`,
	"IFCQUANTITYTIME": `ENTITY IfcQuantityTime
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	TimeValue : IfcTimeMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : NOT(EXISTS(SELF\IfcPhysicalSimpleQuantity.Unit)) OR
   (SELF\IfcPhysicalSimpleQuantity.Unit.UnitType = IfcUnitEnum.TIMEUNIT);
	WR22 : TimeValue >= 0.;
END_ENTITY;`,
	"IFCQUANTITYVOLUME": `ENTITY IfcQuantityVolume
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	VolumeValue : IfcVolumeMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : NOT(EXISTS(SELF\IfcPhysicalSimpleQuantity.Unit)) OR
   (SELF\IfcPhysicalSimpleQuantity.Unit.UnitType = IfcUnitEnum.VOLUMEUNIT);
	WR22 : VolumeValue >= 0.;
END_ENTITY;`,
	"IFCQUANTITYWEIGHT": `ENTITY IfcQuantityWeight
 SUBTYPE OF (IfcPhysicalSimpleQuantity);
	WeightValue : IfcMassMeasure;
	Formula : OPTIONAL IfcLabel;
 WHERE
	WR21 : NOT(EXISTS(SELF\IfcPhysicalSimpleQuantity.Unit)) OR
   (SELF\IfcPhysicalSimpleQuantity.Unit.UnitType = IfcUnitEnum.MASSUNIT);
	WR22 : WeightValue >= 0.;
END_ENTITY;`,
	"IFCRAIL": `ENTITY IfcRail
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcRailTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRailTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCRAILTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCRAILING": `ENTITY IfcRailing
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcRailingTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRailingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCRAILINGTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCRAILINGTYPE": `ENTITY IfcRailingType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcRailingTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcRailingTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailingTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCRAILTYPE": `ENTITY IfcRailType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcRailTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcRailTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCRAILWAY": `ENTITY IfcRailway
 SUBTYPE OF (IfcFacility);
	PredefinedType : OPTIONAL IfcRailwayTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRailwayTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailwayTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCRAILWAYPART": `ENTITY IfcRailwayPart
 SUBTYPE OF (IfcFacilityPart);
	PredefinedType : OPTIONAL IfcRailwayPartTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRailwayPartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRailwayPartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCRAMP": `ENTITY IfcRamp
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcRampTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRampTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRampTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCRAMPTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCRAMPFLIGHT": `ENTITY IfcRampFlight
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcRampFlightTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRampFlightTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRampFlightTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCRAMPFLIGHTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCRAMPFLIGHTTYPE": `ENTITY IfcRampFlightType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcRampFlightTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcRampFlightTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRampFlightTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCRAMPTYPE": `ENTITY IfcRampType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcRampTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcRampTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRampTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCRATIONALBSPLINECURVEWITHKNOTS": `ENTITY IfcRationalBSplineCurveWithKnots
 SUBTYPE OF (IfcBSplineCurveWithKnots);
	WeightsData : LIST [2:?] OF IfcReal;
 DERIVE
	 Weights : ARRAY [0:UpperIndexOnControlPoints] OF IfcReal := IfcListToArray(WeightsData,0,SELF\IfcBSplineCurve.UpperIndexOnControlPoints);
 WHERE
	SameNumOfWeightsAndPoints : SIZEOF(WeightsData) = SIZEOF(SELF\IfcBSplineCurve.ControlPointsList);
	WeightsGreaterZero : IfcCurveWeightsPositive(SELF);
END_ENTITY;`,
	"IFCRATIONALBSPLINESURFACEWITHKNOTS": `ENTITY IfcRationalBSplineSurfaceWithKnots
 SUBTYPE OF (IfcBSplineSurfaceWithKnots);
	WeightsData : LIST [2:?] OF LIST [2:?] OF IfcReal;
 DERIVE
	 Weights : ARRAY [0:UUpper] OF ARRAY [0:VUpper] OF IfcReal := IfcMakeArrayOfArray(WeightsData,0,UUpper,0,VUpper);
 WHERE
	CorrespondingWeightsDataLists : (SIZEOF(WeightsData) = SIZEOF(SELF\IfcBSplineSurface.ControlPointsList))
AND 
(SIZEOF(WeightsData[1]) = SIZEOF(SELF\IfcBSplineSurface.ControlPointsList[1]));
	WeightValuesGreaterZero : IfcSurfaceWeightsPositive(SELF);
END_ENTITY;`,
	"IFCRECTANGLEHOLLOWPROFILEDEF": `ENTITY IfcRectangleHollowProfileDef
 SUBTYPE OF (IfcRectangleProfileDef);
	WallThickness : IfcPositiveLengthMeasure;
	InnerFilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	OuterFilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
 WHERE
	ValidInnerRadius : NOT(EXISTS(InnerFilletRadius)) OR 
((InnerFilletRadius <= (SELF\IfcRectangleProfileDef.XDim/2. - WallThickness)) AND 
 (InnerFilletRadius <= (SELF\IfcRectangleProfileDef.YDim/2. - WallThickness)))
;
	ValidOuterRadius : NOT(EXISTS(OuterFilletRadius)) OR 
((OuterFilletRadius <= (SELF\IfcRectangleProfileDef.XDim/2.)) AND 
 (OuterFilletRadius <= (SELF\IfcRectangleProfileDef.YDim/2.)));
	ValidWallThickness : (WallThickness < (SELF\IfcRectangleProfileDef.XDim/2.)) AND 
(WallThickness < (SELF\IfcRectangleProfileDef.YDim/2.));
END_ENTITY;`,
	"IFCRECTANGLEPROFILEDEF": `ENTITY IfcRectangleProfileDef
 SUPERTYPE OF (ONEOF
	(IfcRectangleHollowProfileDef
	,IfcRoundedRectangleProfileDef))
 SUBTYPE OF (IfcParameterizedProfileDef);
	XDim : IfcPositiveLengthMeasure;
	YDim : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCRECTANGULARPYRAMID": `ENTITY IfcRectangularPyramid
 SUBTYPE OF (IfcCsgPrimitive3D);
	XLength : IfcPositiveLengthMeasure;
	YLength : IfcPositiveLengthMeasure;
	Height : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCRECTANGULARTRIMMEDSURFACE": `ENTITY IfcRectangularTrimmedSurface
 SUBTYPE OF (IfcBoundedSurface);
	BasisSurface : IfcSurface;
	U1 : IfcParameterValue;
	V1 : IfcParameterValue;
	U2 : IfcParameterValue;
	V2 : IfcParameterValue;
	Usense : IfcBoolean;
	Vsense : IfcBoolean;
 WHERE
	U1AndU2Different : U1 <> U2;
	UsenseCompatible : (('IFC4X3_DEV_73740fe4.IFCELEMENTARYSURFACE' IN TYPEOF(BasisSurface)) AND
         (NOT ('IFC4X3_DEV_73740fe4.IFCPLANE' IN TYPEOF(BasisSurface)))) OR
         ('IFC4X3_DEV_73740fe4.IFCSURFACEOFREVOLUTION' IN TYPEOF(BasisSurface)) OR
         (Usense = (U2 > U1));
	V1AndV2Different : V1 <> V2;
	VsenseCompatible : Vsense = (V2 > V1);
END_ENTITY;`,
	"IFCRECURRENCEPATTERN": `ENTITY IfcRecurrencePattern;
	RecurrenceType : IfcRecurrenceTypeEnum;
	DayComponent : OPTIONAL SET [1:?] OF IfcDayInMonthNumber;
	WeekdayComponent : OPTIONAL SET [1:?] OF IfcDayInWeekNumber;
	MonthComponent : OPTIONAL SET [1:?] OF IfcMonthInYearNumber;
	Position : OPTIONAL IfcInteger;
	Interval : OPTIONAL IfcInteger;
	Occurrences : OPTIONAL IfcInteger;
	TimePeriods : OPTIONAL LIST [1:?] OF IfcTimePeriod;
END_ENTITY;`,
	"IFCREFERENCE": `ENTITY IfcReference;
	TypeIdentifier : OPTIONAL IfcIdentifier;
	AttributeIdentifier : OPTIONAL IfcIdentifier;
	InstanceName : OPTIONAL IfcLabel;
	ListPositions : OPTIONAL LIST [1:?] OF IfcInteger;
	InnerReference : OPTIONAL IfcReference;
END_ENTITY;`,
	"IFCREFERENT": `ENTITY IfcReferent
 SUBTYPE OF (IfcPositioningElement);
	PredefinedType : OPTIONAL IfcReferentTypeEnum;
END_ENTITY;`,
	"IFCREGULARTIMESERIES": `ENTITY IfcRegularTimeSeries
 SUBTYPE OF (IfcTimeSeries);
	TimeStep : IfcTimeMeasure;
	Values : LIST [1:?] OF IfcTimeSeriesValue;
END_ENTITY;`,
	"IFCREINFORCEDSOIL": `ENTITY IfcReinforcedSoil
 SUBTYPE OF (IfcEarthworksElement);
	PredefinedType : OPTIONAL IfcReinforcedSoilTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcReinforcedSoilTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcReinforcedSoilTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCREINFORCEMENTBARPROPERTIES": `ENTITY IfcReinforcementBarProperties
 SUBTYPE OF (IfcPreDefinedProperties);
	TotalCrossSectionArea : IfcAreaMeasure;
	SteelGrade : IfcLabel;
	BarSurface : OPTIONAL IfcReinforcingBarSurfaceEnum;
	EffectiveDepth : OPTIONAL IfcLengthMeasure;
	NominalBarDiameter : OPTIONAL IfcPositiveLengthMeasure;
	BarCount : OPTIONAL IfcCountMeasure;
END_ENTITY;`,
	"IFCREINFORCEMENTDEFINITIONPROPERTIES": `ENTITY IfcReinforcementDefinitionProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	DefinitionType : OPTIONAL IfcLabel;
	ReinforcementSectionDefinitions : LIST [1:?] OF IfcSectionReinforcementProperties;
END_ENTITY;`,
	"IFCREINFORCINGBAR": `ENTITY IfcReinforcingBar
 SUBTYPE OF (IfcReinforcingElement);
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	CrossSectionArea : OPTIONAL IfcAreaMeasure;
	BarLength : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcReinforcingBarTypeEnum;
	BarSurface : OPTIONAL IfcReinforcingBarSurfaceEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcReinforcingBarTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcReinforcingBarTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCREINFORCINGBARTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCREINFORCINGBARTYPE": `ENTITY IfcReinforcingBarType
 SUBTYPE OF (IfcReinforcingElementType);
	PredefinedType : IfcReinforcingBarTypeEnum;
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	CrossSectionArea : OPTIONAL IfcAreaMeasure;
	BarLength : OPTIONAL IfcPositiveLengthMeasure;
	BarSurface : OPTIONAL IfcReinforcingBarSurfaceEnum;
	BendingShapeCode : OPTIONAL IfcLabel;
	BendingParameters : OPTIONAL LIST [1:?] OF IfcBendingParameterSelect;
 WHERE
	BendingShapeCodeProvided : NOT EXISTS(BendingParameters) OR EXISTS(BendingShapeCode);
	CorrectPredefinedType : (PredefinedType <> IfcReinforcingBarTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcReinforcingBarTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCREINFORCINGELEMENT": `ENTITY IfcReinforcingElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcReinforcingBar
	,IfcReinforcingMesh
	,IfcTendon
	,IfcTendonAnchor
	,IfcTendonConduit))
 SUBTYPE OF (IfcElementComponent);
	SteelGrade : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCREINFORCINGELEMENTTYPE": `ENTITY IfcReinforcingElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcReinforcingBarType
	,IfcReinforcingMeshType
	,IfcTendonAnchorType
	,IfcTendonConduitType
	,IfcTendonType))
 SUBTYPE OF (IfcElementComponentType);
END_ENTITY;`,
	"IFCREINFORCINGMESH": `ENTITY IfcReinforcingMesh
 SUBTYPE OF (IfcReinforcingElement);
	MeshLength : OPTIONAL IfcPositiveLengthMeasure;
	MeshWidth : OPTIONAL IfcPositiveLengthMeasure;
	LongitudinalBarNominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	TransverseBarNominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	LongitudinalBarCrossSectionArea : OPTIONAL IfcAreaMeasure;
	TransverseBarCrossSectionArea : OPTIONAL IfcAreaMeasure;
	LongitudinalBarSpacing : OPTIONAL IfcPositiveLengthMeasure;
	TransverseBarSpacing : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcReinforcingMeshTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcReinforcingMeshTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcReinforcingMeshTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCREINFORCINGMESHTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCREINFORCINGMESHTYPE": `ENTITY IfcReinforcingMeshType
 SUBTYPE OF (IfcReinforcingElementType);
	PredefinedType : IfcReinforcingMeshTypeEnum;
	MeshLength : OPTIONAL IfcPositiveLengthMeasure;
	MeshWidth : OPTIONAL IfcPositiveLengthMeasure;
	LongitudinalBarNominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	TransverseBarNominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	LongitudinalBarCrossSectionArea : OPTIONAL IfcAreaMeasure;
	TransverseBarCrossSectionArea : OPTIONAL IfcAreaMeasure;
	LongitudinalBarSpacing : OPTIONAL IfcPositiveLengthMeasure;
	TransverseBarSpacing : OPTIONAL IfcPositiveLengthMeasure;
	BendingShapeCode : OPTIONAL IfcLabel;
	BendingParameters : OPTIONAL LIST [1:?] OF IfcBendingParameterSelect;
 WHERE
	BendingShapeCodeProvided : NOT EXISTS(BendingParameters) OR EXISTS(BendingShapeCode);
	CorrectPredefinedType : (PredefinedType <> IfcReinforcingMeshTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcReinforcingMeshTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCRELADHERESTOELEMENT": `ENTITY IfcRelAdheresToElement
 SUBTYPE OF (IfcRelDecomposes);
	RelatingElement : IfcElement;
	RelatedSurfaceFeatures : SET [1:?] OF IfcSurfaceFeature;
END_ENTITY;`,
	"IFCRELAGGREGATES": `ENTITY IfcRelAggregates
 SUBTYPE OF (IfcRelDecomposes);
	RelatingObject : IfcObjectDefinition;
	RelatedObjects : SET [1:?] OF IfcObjectDefinition;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* RelatedObjects | RelatingObject :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNS": `ENTITY IfcRelAssigns
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelAssignsToActor
	,IfcRelAssignsToControl
	,IfcRelAssignsToGroup
	,IfcRelAssignsToProcess
	,IfcRelAssignsToProduct
	,IfcRelAssignsToResource))
 SUBTYPE OF (IfcRelationship);
	RelatedObjects : SET [1:?] OF IfcObjectDefinition;
	RelatedObjectsType : OPTIONAL IfcStrippedOptional;
END_ENTITY;`,
	"IFCRELASSIGNSTOACTOR": `ENTITY IfcRelAssignsToActor
 SUBTYPE OF (IfcRelAssigns);
	RelatingActor : IfcActor;
	ActingRole : OPTIONAL IfcActorRole;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingActor :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNSTOCONTROL": `ENTITY IfcRelAssignsToControl
 SUBTYPE OF (IfcRelAssigns);
	RelatingControl : IfcControl;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingControl :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNSTOGROUP": `ENTITY IfcRelAssignsToGroup
 SUPERTYPE OF (ONEOF
	(IfcRelAssignsToGroupByFactor))
 SUBTYPE OF (IfcRelAssigns);
	RelatingGroup : IfcGroup;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingGroup :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNSTOGROUPBYFACTOR": `ENTITY IfcRelAssignsToGroupByFactor
 SUBTYPE OF (IfcRelAssignsToGroup);
	Factor : IfcRatioMeasure;
END_ENTITY;`,
	"IFCRELASSIGNSTOPROCESS": `ENTITY IfcRelAssignsToProcess
 SUBTYPE OF (IfcRelAssigns);
	RelatingProcess : IfcProcessSelect;
	QuantityInProcess : OPTIONAL IfcMeasureWithUnit;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingProcess :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNSTOPRODUCT": `ENTITY IfcRelAssignsToProduct
 SUBTYPE OF (IfcRelAssigns);
	RelatingProduct : IfcProductSelect;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingProduct :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSIGNSTORESOURCE": `ENTITY IfcRelAssignsToResource
 SUBTYPE OF (IfcRelAssigns);
	RelatingResource : IfcResourceSelect;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* SELF\IfcRelAssigns.RelatedObjects | RelatingResource :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELASSOCIATES": `ENTITY IfcRelAssociates
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelAssociatesApproval
	,IfcRelAssociatesClassification
	,IfcRelAssociatesConstraint
	,IfcRelAssociatesDocument
	,IfcRelAssociatesLibrary
	,IfcRelAssociatesMaterial
	,IfcRelAssociatesProfileDef))
 SUBTYPE OF (IfcRelationship);
	RelatedObjects : SET [1:?] OF IfcDefinitionSelect;
END_ENTITY;`,
	"IFCRELASSOCIATESAPPROVAL": `ENTITY IfcRelAssociatesApproval
 SUBTYPE OF (IfcRelAssociates);
	RelatingApproval : IfcApproval;
END_ENTITY;`,
	"IFCRELASSOCIATESCLASSIFICATION": `ENTITY IfcRelAssociatesClassification
 SUBTYPE OF (IfcRelAssociates);
	RelatingClassification : IfcClassificationSelect;
END_ENTITY;`,
	"IFCRELASSOCIATESCONSTRAINT": `ENTITY IfcRelAssociatesConstraint
 SUBTYPE OF (IfcRelAssociates);
	Intent : OPTIONAL IfcLabel;
	RelatingConstraint : IfcConstraint;
END_ENTITY;`,
	"IFCRELASSOCIATESDOCUMENT": `ENTITY IfcRelAssociatesDocument
 SUBTYPE OF (IfcRelAssociates);
	RelatingDocument : IfcDocumentSelect;
END_ENTITY;`,
	"IFCRELASSOCIATESLIBRARY": `ENTITY IfcRelAssociatesLibrary
 SUBTYPE OF (IfcRelAssociates);
	RelatingLibrary : IfcLibrarySelect;
END_ENTITY;`,
	"IFCRELASSOCIATESMATERIAL": `ENTITY IfcRelAssociatesMaterial
 SUBTYPE OF (IfcRelAssociates);
	RelatingMaterial : IfcMaterialSelect;
 WHERE
	AllowedElements : SIZEOF(QUERY(temp <* SELF\IfcRelAssociates.RelatedObjects | (
  SIZEOF(TYPEOF(temp) * [   
    'IFC4X3_DEV_73740fe4.IFCELEMENT', 
    'IFC4X3_DEV_73740fe4.IFCELEMENTTYPE',
    'IFC4X3_DEV_73740fe4.IFCSTRUCTURALMEMBER',
    'IFC4X3_DEV_73740fe4.IFCPORT']) = 0) 
)) = 0;
	NoVoidElement : SIZEOF(QUERY(temp <* SELF\IfcRelAssociates.RelatedObjects | 
  ('IFC4X3_DEV_73740fe4.IFCFEATUREELEMENTSUBTRACTION' IN TYPEOF(temp)) OR 
  ('IFC4X3_DEV_73740fe4.IFCVIRTUALELEMENT' IN TYPEOF(temp)) 
)) = 0;
END_ENTITY;`,
	"IFCRELASSOCIATESPROFILEDEF": `ENTITY IfcRelAssociatesProfileDef
 SUBTYPE OF (IfcRelAssociates);
	RelatingProfileDef : IfcProfileDef;
END_ENTITY;`,
	"IFCRELATIONSHIP": `ENTITY IfcRelationship
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelAssigns
	,IfcRelAssociates
	,IfcRelConnects
	,IfcRelDeclares
	,IfcRelDecomposes
	,IfcRelDefines))
 SUBTYPE OF (IfcRoot);
END_ENTITY;`,
	"IFCRELCONNECTS": `ENTITY IfcRelConnects
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelConnectsElements
	,IfcRelConnectsPortToElement
	,IfcRelConnectsPorts
	,IfcRelConnectsStructuralActivity
	,IfcRelConnectsStructuralMember
	,IfcRelContainedInSpatialStructure
	,IfcRelCoversBldgElements
	,IfcRelCoversSpaces
	,IfcRelFillsElement
	,IfcRelFlowControlElements
	,IfcRelInterferesElements
	,IfcRelPositions
	,IfcRelReferencedInSpatialStructure
	,IfcRelSequence
	,IfcRelServicesBuildings
	,IfcRelSpaceBoundary))
 SUBTYPE OF (IfcRelationship);
END_ENTITY;`,
	"IFCRELCONNECTSELEMENTS": `ENTITY IfcRelConnectsElements
 SUPERTYPE OF (ONEOF
	(IfcRelConnectsPathElements
	,IfcRelConnectsWithRealizingElements))
 SUBTYPE OF (IfcRelConnects);
	ConnectionGeometry : OPTIONAL IfcConnectionGeometry;
	RelatingElement : IfcElement;
	RelatedElement : IfcElement;
 WHERE
	NoSelfReference : RelatingElement :<>: RelatedElement;
END_ENTITY;`,
	"IFCRELCONNECTSPATHELEMENTS": `ENTITY IfcRelConnectsPathElements
 SUBTYPE OF (IfcRelConnectsElements);
	RelatingPriorities : LIST [0:?] OF IfcInteger;
	RelatedPriorities : LIST [0:?] OF IfcInteger;
	RelatedConnectionType : IfcConnectionTypeEnum;
	RelatingConnectionType : IfcConnectionTypeEnum;
 WHERE
	NormalizedRelatedPriorities : (SIZEOF(RelatedPriorities) = 0)
OR
(SIZEOF (QUERY (temp <* RelatedPriorities
  | {0 <= temp <= 100}
  )) = SIZEOF(RelatedPriorities));
	NormalizedRelatingPriorities : (SIZEOF(RelatingPriorities) = 0)
OR
(SIZEOF (QUERY (temp <* RelatingPriorities
  | {0 <= temp <= 100}
  )) = SIZEOF(RelatingPriorities));
END_ENTITY;`,
	"IFCRELCONNECTSPORTS": `ENTITY IfcRelConnectsPorts
 SUBTYPE OF (IfcRelConnects);
	RelatingPort : IfcPort;
	RelatedPort : IfcPort;
	RealizingElement : OPTIONAL IfcElement;
 WHERE
	NoSelfReference : RelatingPort :<>: RelatedPort;
END_ENTITY;`,
	"IFCRELCONNECTSPORTTOELEMENT": `ENTITY IfcRelConnectsPortToElement
 SUBTYPE OF (IfcRelConnects);
	RelatingPort : IfcPort;
	RelatedElement : IfcDistributionElement;
END_ENTITY;`,
	"IFCRELCONNECTSSTRUCTURALACTIVITY": `ENTITY IfcRelConnectsStructuralActivity
 SUBTYPE OF (IfcRelConnects);
	RelatingElement : IfcStructuralActivityAssignmentSelect;
	RelatedStructuralActivity : IfcStructuralActivity;
END_ENTITY;`,
	"IFCRELCONNECTSSTRUCTURALMEMBER": `ENTITY IfcRelConnectsStructuralMember
 SUPERTYPE OF (ONEOF
	(IfcRelConnectsWithEccentricity))
 SUBTYPE OF (IfcRelConnects);
	RelatingStructuralMember : IfcStructuralMember;
	RelatedStructuralConnection : IfcStructuralConnection;
	AppliedCondition : OPTIONAL IfcBoundaryCondition;
	AdditionalConditions : OPTIONAL IfcStructuralConnectionCondition;
	SupportedLength : OPTIONAL IfcLengthMeasure;
	ConditionCoordinateSystem : OPTIONAL IfcAxis2Placement3D;
END_ENTITY;`,
	"IFCRELCONNECTSWITHECCENTRICITY": `ENTITY IfcRelConnectsWithEccentricity
 SUBTYPE OF (IfcRelConnectsStructuralMember);
	ConnectionConstraint : IfcConnectionGeometry;
END_ENTITY;`,
	"IFCRELCONNECTSWITHREALIZINGELEMENTS": `ENTITY IfcRelConnectsWithRealizingElements
 SUBTYPE OF (IfcRelConnectsElements);
	RealizingElements : SET [1:?] OF IfcElement;
	ConnectionType : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCRELCONTAINEDINSPATIALSTRUCTURE": `ENTITY IfcRelContainedInSpatialStructure
 SUBTYPE OF (IfcRelConnects);
	RelatedElements : SET [1:?] OF IfcProduct;
	RelatingStructure : IfcSpatialElement;
 WHERE
	WR31 : SIZEOF(QUERY(temp <* RelatedElements | 'IFC4X3_DEV_73740fe4.IFCSPATIALSTRUCTUREELEMENT' IN TYPEOF(temp))) = 0;
END_ENTITY;`,
	"IFCRELCOVERSBLDGELEMENTS": `ENTITY IfcRelCoversBldgElements
 SUBTYPE OF (IfcRelConnects);
	RelatingBuildingElement : IfcElement;
	RelatedCoverings : SET [1:?] OF IfcCovering;
END_ENTITY;`,
	"IFCRELCOVERSSPACES": `ENTITY IfcRelCoversSpaces
 SUBTYPE OF (IfcRelConnects);
	RelatingSpace : IfcSpace;
	RelatedCoverings : SET [1:?] OF IfcCovering;
END_ENTITY;`,
	"IFCRELDECLARES": `ENTITY IfcRelDeclares
 SUBTYPE OF (IfcRelationship);
	RelatingContext : IfcContext;
	RelatedDefinitions : SET [1:?] OF IfcDefinitionSelect;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* RelatedDefinitions | RelatingContext :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELDECOMPOSES": `ENTITY IfcRelDecomposes
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelAdheresToElement
	,IfcRelAggregates
	,IfcRelNests
	,IfcRelProjectsElement
	,IfcRelVoidsElement))
 SUBTYPE OF (IfcRelationship);
END_ENTITY;`,
	"IFCRELDEFINES": `ENTITY IfcRelDefines
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcRelDefinesByObject
	,IfcRelDefinesByProperties
	,IfcRelDefinesByTemplate
	,IfcRelDefinesByType))
 SUBTYPE OF (IfcRelationship);
END_ENTITY;`,
	"IFCRELDEFINESBYOBJECT": `ENTITY IfcRelDefinesByObject
 SUBTYPE OF (IfcRelDefines);
	RelatedObjects : SET [1:?] OF IfcObject;
	RelatingObject : IfcObject;
END_ENTITY;`,
	"IFCRELDEFINESBYPROPERTIES": `ENTITY IfcRelDefinesByProperties
 SUBTYPE OF (IfcRelDefines);
	RelatedObjects : SET [1:?] OF IfcObjectDefinition;
	RelatingPropertyDefinition : IfcPropertySetDefinitionSelect;
 WHERE
	NoRelatedTypeObject : SIZEOF(QUERY(Types <* SELF\IfcRelDefinesByProperties.RelatedObjects |  'IFC4X3_DEV_73740fe4.IFCTYPEOBJECT' IN TYPEOF(Types))) = 0;
END_ENTITY;`,
	"IFCRELDEFINESBYTEMPLATE": `ENTITY IfcRelDefinesByTemplate
 SUBTYPE OF (IfcRelDefines);
	RelatedPropertySets : SET [1:?] OF IfcPropertySetDefinition;
	RelatingTemplate : IfcPropertySetTemplate;
END_ENTITY;`,
	"IFCRELDEFINESBYTYPE": `ENTITY IfcRelDefinesByType
 SUBTYPE OF (IfcRelDefines);
	RelatedObjects : SET [1:?] OF IfcObject;
	RelatingType : IfcTypeObject;
END_ENTITY;`,
	"IFCRELFILLSELEMENT": `ENTITY IfcRelFillsElement
 SUBTYPE OF (IfcRelConnects);
	RelatingOpeningElement : IfcOpeningElement;
	RelatedBuildingElement : IfcElement;
END_ENTITY;`,
	"IFCRELFLOWCONTROLELEMENTS": `ENTITY IfcRelFlowControlElements
 SUBTYPE OF (IfcRelConnects);
	RelatedControlElements : SET [1:?] OF IfcDistributionControlElement;
	RelatingFlowElement : IfcDistributionFlowElement;
END_ENTITY;`,
	"IFCRELINTERFERESELEMENTS": `ENTITY IfcRelInterferesElements
 SUBTYPE OF (IfcRelConnects);
	RelatingElement : IfcInterferenceSelect;
	RelatedElement : IfcInterferenceSelect;
	InterferenceGeometry : OPTIONAL IfcConnectionGeometry;
	InterferenceType : OPTIONAL IfcIdentifier;
	ImpliedOrder : IfcLogical;
	InterferenceSpace : OPTIONAL IfcSpatialZone;
 WHERE
	NoSelfReference : RelatingElement :<>: RelatedElement;
END_ENTITY;`,
	"IFCRELNESTS": `ENTITY IfcRelNests
 SUBTYPE OF (IfcRelDecomposes);
	RelatingObject : IfcObjectDefinition;
	RelatedObjects : LIST [1:?] OF IfcObjectDefinition;
 WHERE
	NoSelfReference : SIZEOF(QUERY(Temp <* RelatedObjects | RelatingObject :=: Temp)) = 0;
END_ENTITY;`,
	"IFCRELPOSITIONS": `ENTITY IfcRelPositions
 SUBTYPE OF (IfcRelConnects);
	RelatingPositioningElement : IfcPositioningElement;
	RelatedProducts : SET [1:?] OF IfcProduct;
 WHERE
	NoSelfReference : SIZEOF(QUERY(temp <* RelatedProducts | RelatingPositioningElement :=: temp)) = 0;
END_ENTITY;`,
	"IFCRELPROJECTSELEMENT": `ENTITY IfcRelProjectsElement
 SUBTYPE OF (IfcRelDecomposes);
	RelatingElement : IfcElement;
	RelatedFeatureElement : IfcFeatureElementAddition;
END_ENTITY;`,
	"IFCRELREFERENCEDINSPATIALSTRUCTURE": `ENTITY IfcRelReferencedInSpatialStructure
 SUBTYPE OF (IfcRelConnects);
	RelatedElements : SET [1:?] OF IfcSpatialReferenceSelect;
	RelatingStructure : IfcSpatialElement;
 WHERE
	AllowedRelatedElements : SIZEOF(QUERY(temp <* RelatedElements | ('IFC4X3_DEV_73740fe4.IFCSPATIALSTRUCTUREELEMENT' IN TYPEOF(temp)) AND (NOT ('IFC4X3_DEV_73740fe4.IFCSPACE' IN TYPEOF(temp)))
)) = 0;
END_ENTITY;`,
	"IFCRELSEQUENCE": `ENTITY IfcRelSequence
 SUBTYPE OF (IfcRelConnects);
	RelatingProcess : IfcProcess;
	RelatedProcess : IfcProcess;
	TimeLag : OPTIONAL IfcLagTime;
	SequenceType : OPTIONAL IfcSequenceEnum;
	UserDefinedSequenceType : OPTIONAL IfcLabel;
 WHERE
	AvoidInconsistentSequence : RelatingProcess :<>: RelatedProcess;
	CorrectSequenceType : (SequenceType <> IfcSequenceEnum.USERDEFINED) OR ((SequenceType = IfcSequenceEnum.USERDEFINED) AND EXISTS(UserDefinedSequenceType));
END_ENTITY;`,
	"IFCRELSERVICESBUILDINGS": `ENTITY IfcRelServicesBuildings
 SUBTYPE OF (IfcRelConnects);
	RelatingSystem : IfcSystem;
	RelatedBuildings : SET [1:?] OF IfcSpatialElement;
END_ENTITY;`,
	"IFCRELSPACEBOUNDARY": `ENTITY IfcRelSpaceBoundary
 SUPERTYPE OF (ONEOF
	(IfcRelSpaceBoundary1stLevel))
 SUBTYPE OF (IfcRelConnects);
	RelatingSpace : IfcSpaceBoundarySelect;
	RelatedBuildingElement : IfcElement;
	ConnectionGeometry : OPTIONAL IfcConnectionGeometry;
	PhysicalOrVirtualBoundary : IfcPhysicalOrVirtualEnum;
	InternalOrExternalBoundary : IfcInternalOrExternalEnum;
 WHERE
	CorrectPhysOrVirt : ((PhysicalOrVirtualBoundary = IfcPhysicalOrVirtualEnum.Physical) 
  AND (NOT('IFC4X3_DEV_73740fe4.IFCVIRTUALELEMENT' IN TYPEOF(RelatedBuildingElement)))) 
OR 
((PhysicalOrVirtualBoundary = IfcPhysicalOrVirtualEnum.Virtual) 
  AND (('IFC4X3_DEV_73740fe4.IFCVIRTUALELEMENT' IN TYPEOF(RelatedBuildingElement)) 
  OR   ('IFC4X3_DEV_73740fe4.IFCOPENINGELEMENT' IN TYPEOF(RelatedBuildingElement)))) 
OR 
(PhysicalOrVirtualBoundary = IfcPhysicalOrVirtualEnum.NotDefined);
END_ENTITY;`,
	"IFCRELSPACEBOUNDARY1STLEVEL": `ENTITY IfcRelSpaceBoundary1stLevel
 SUPERTYPE OF (ONEOF
	(IfcRelSpaceBoundary2ndLevel))
 SUBTYPE OF (IfcRelSpaceBoundary);
	ParentBoundary : OPTIONAL IfcRelSpaceBoundary1stLevel;
 INVERSE
	InnerBoundaries : SET [0:?] OF IfcRelSpaceBoundary1stLevel FOR ParentBoundary;
END_ENTITY;`,
	"IFCRELSPACEBOUNDARY2NDLEVEL": `ENTITY IfcRelSpaceBoundary2ndLevel
 SUBTYPE OF (IfcRelSpaceBoundary1stLevel);
	CorrespondingBoundary : OPTIONAL IfcRelSpaceBoundary2ndLevel;
 INVERSE
	Corresponds : SET [0:1] OF IfcRelSpaceBoundary2ndLevel FOR CorrespondingBoundary;
END_ENTITY;`,
	"IFCRELVOIDSELEMENT": `ENTITY IfcRelVoidsElement
 SUBTYPE OF (IfcRelDecomposes);
	RelatingBuildingElement : IfcElement;
	RelatedOpeningElement : IfcFeatureElementSubtraction;
END_ENTITY;`,
	"IFCREPARAMETRISEDCOMPOSITECURVESEGMENT": `ENTITY IfcReparametrisedCompositeCurveSegment
 SUBTYPE OF (IfcCompositeCurveSegment);
	ParamLength : IfcParameterValue;
 WHERE
	PositiveLengthParameter : ParamLength > 0.0;
END_ENTITY;`,
	"IFCREPRESENTATION": `ENTITY IfcRepresentation
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcShapeModel
	,IfcStyleModel));
	ContextOfItems : IfcRepresentationContext;
	RepresentationIdentifier : OPTIONAL IfcLabel;
	RepresentationType : OPTIONAL IfcLabel;
	Items : SET [1:?] OF IfcRepresentationItem;
 INVERSE
	RepresentationMap : SET [0:1] OF IfcRepresentationMap FOR MappedRepresentation;
	LayerAssignments : SET [0:?] OF IfcPresentationLayerAssignment FOR AssignedItems;
	OfProductRepresentation : SET [0:?] OF IfcProductRepresentation FOR Representations;
END_ENTITY;`,
	"IFCREPRESENTATIONCONTEXT": `ENTITY IfcRepresentationContext
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGeometricRepresentationContext));
	ContextIdentifier : OPTIONAL IfcLabel;
	ContextType : OPTIONAL IfcLabel;
 INVERSE
	RepresentationsInContext : SET [0:?] OF IfcRepresentation FOR ContextOfItems;
END_ENTITY;`,
	"IFCREPRESENTATIONITEM": `ENTITY IfcRepresentationItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcGeometricRepresentationItem
	,IfcMappedItem
	,IfcStyledItem
	,IfcTopologicalRepresentationItem));
 INVERSE
	LayerAssignment : SET [0:1] OF IfcPresentationLayerAssignment FOR AssignedItems;
	StyledByItem : SET [0:1] OF IfcStyledItem FOR Item;
END_ENTITY;`,
	"IFCREPRESENTATIONMAP": `ENTITY IfcRepresentationMap;
	MappingOrigin : IfcAxis2Placement;
	MappedRepresentation : IfcRepresentation;
 INVERSE
	HasShapeAspects : SET [0:?] OF IfcShapeAspect FOR PartOfProductDefinitionShape;
	MapUsage : SET [0:?] OF IfcMappedItem FOR MappingSource;
 WHERE
	ApplicableMappedRepr : 'IFC4X3_DEV_73740fe4.IFCSHAPEMODEL' IN TYPEOF(MappedRepresentation);
END_ENTITY;`,
	"IFCRESOURCE": `ENTITY IfcResource
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConstructionResource))
 SUBTYPE OF (IfcObject);
	Identification : OPTIONAL IfcIdentifier;
	LongDescription : OPTIONAL IfcText;
 INVERSE
	ResourceOf : SET [0:?] OF IfcRelAssignsToResource FOR RelatingResource;
END_ENTITY;`,
	"IFCRESOURCEAPPROVALRELATIONSHIP": `ENTITY IfcResourceApprovalRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatedResourceObjects : SET [1:?] OF IfcResourceObjectSelect;
	RelatingApproval : IfcApproval;
END_ENTITY;`,
	"IFCRESOURCECONSTRAINTRELATIONSHIP": `ENTITY IfcResourceConstraintRelationship
 SUBTYPE OF (IfcResourceLevelRelationship);
	RelatingConstraint : IfcConstraint;
	RelatedResourceObjects : SET [1:?] OF IfcResourceObjectSelect;
END_ENTITY;`,
	"IFCRESOURCELEVELRELATIONSHIP": `ENTITY IfcResourceLevelRelationship
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcApprovalRelationship
	,IfcCurrencyRelationship
	,IfcDocumentInformationRelationship
	,IfcExternalReferenceRelationship
	,IfcMaterialRelationship
	,IfcOrganizationRelationship
	,IfcPropertyDependencyRelationship
	,IfcResourceApprovalRelationship
	,IfcResourceConstraintRelationship));
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
END_ENTITY;`,
	"IFCRESOURCETIME": `ENTITY IfcResourceTime
 SUBTYPE OF (IfcSchedulingTime);
	ScheduleWork : OPTIONAL IfcDuration;
	ScheduleUsage : OPTIONAL IfcPositiveRatioMeasure;
	ScheduleStart : OPTIONAL IfcDateTime;
	ScheduleFinish : OPTIONAL IfcDateTime;
	ScheduleContour : OPTIONAL IfcLabel;
	LevelingDelay : OPTIONAL IfcDuration;
	IsOverAllocated : OPTIONAL IfcBoolean;
	StatusTime : OPTIONAL IfcDateTime;
	ActualWork : OPTIONAL IfcDuration;
	ActualUsage : OPTIONAL IfcPositiveRatioMeasure;
	ActualStart : OPTIONAL IfcDateTime;
	ActualFinish : OPTIONAL IfcDateTime;
	RemainingWork : OPTIONAL IfcDuration;
	RemainingUsage : OPTIONAL IfcPositiveRatioMeasure;
	Completion : OPTIONAL IfcPositiveRatioMeasure;
END_ENTITY;`,
	"IFCREVOLVEDAREASOLID": `ENTITY IfcRevolvedAreaSolid
 SUPERTYPE OF (ONEOF
	(IfcRevolvedAreaSolidTapered))
 SUBTYPE OF (IfcSweptAreaSolid);
	Axis : IfcAxis1Placement;
	Angle : IfcPlaneAngleMeasure;
 DERIVE
	 AxisLine : IfcLine := IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcCurve() || IfcLine(Axis.Location, 
  IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcVector(Axis.Z,1.0));
 WHERE
	AxisDirectionInXY : Axis.Z.DirectionRatios[3] = 0.0;
	AxisStartInXY : ('IFC4X3_DEV_73740fe4.IFCCARTESIANPOINT' IN TYPEOF(Axis.Location)) AND (Axis.Location\IfcCartesianPoint.Coordinates[3] = 0.0);
END_ENTITY;`,
	"IFCREVOLVEDAREASOLIDTAPERED": `ENTITY IfcRevolvedAreaSolidTapered
 SUBTYPE OF (IfcRevolvedAreaSolid);
	EndSweptArea : IfcProfileDef;
 WHERE
	CorrectProfileAssignment : IfcTaperedSweptAreaProfiles(SELF\IfcSweptAreaSolid.SweptArea, SELF.EndSweptArea);
END_ENTITY;`,
	"IFCRIGHTCIRCULARCONE": `ENTITY IfcRightCircularCone
 SUBTYPE OF (IfcCsgPrimitive3D);
	Height : IfcPositiveLengthMeasure;
	BottomRadius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCRIGHTCIRCULARCYLINDER": `ENTITY IfcRightCircularCylinder
 SUBTYPE OF (IfcCsgPrimitive3D);
	Height : IfcPositiveLengthMeasure;
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCRIGIDOPERATION": `ENTITY IfcRigidOperation
 SUBTYPE OF (IfcCoordinateOperation);
	FirstCoordinate : IfcMeasureValue;
	SecondCoordinate : IfcMeasureValue;
	Height : OPTIONAL IfcLengthMeasure;
 WHERE
	SameCoordinateType : (('IFC4X3_DEV_73740fe4.IFCLENGTHMEASURE' IN TYPEOF(FirstCoordinate)) AND ('IFC4X3_DEV_73740fe4.IFCLENGTHMEASURE' IN TYPEOF(SecondCoordinate))) OR (('IFC4X3_DEV_73740fe4.IFCPLANEANGLEMEASURE' IN TYPEOF(FirstCoordinate)) AND ('IFC4X3_DEV_73740fe4.IFCPLANEANGLEMEASURE' IN TYPEOF(SecondCoordinate)));
END_ENTITY;`,
	"IFCROAD": `ENTITY IfcRoad
 SUBTYPE OF (IfcFacility);
	PredefinedType : OPTIONAL IfcRoadTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRoadTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRoadTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCROADPART": `ENTITY IfcRoadPart
 SUBTYPE OF (IfcFacilityPart);
	PredefinedType : OPTIONAL IfcRoadPartTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRoadPartTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRoadPartTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCROOF": `ENTITY IfcRoof
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcRoofTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcRoofTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRoofTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCROOFTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCROOFTYPE": `ENTITY IfcRoofType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcRoofTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcRoofTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcRoofTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCROOT": `ENTITY IfcRoot
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcObjectDefinition
	,IfcPropertyDefinition
	,IfcRelationship));
	GlobalId : IfcGloballyUniqueId;
	OwnerHistory : OPTIONAL IfcOwnerHistory;
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
 UNIQUE
	UR1 : GlobalId;
END_ENTITY;`,
	"IFCROUNDEDRECTANGLEPROFILEDEF": `ENTITY IfcRoundedRectangleProfileDef
 SUBTYPE OF (IfcRectangleProfileDef);
	RoundingRadius : IfcPositiveLengthMeasure;
 WHERE
	ValidRadius : ((RoundingRadius <= (SELF\IfcRectangleProfileDef.XDim/2.)) AND 
 (RoundingRadius <= (SELF\IfcRectangleProfileDef.YDim/2.)));
END_ENTITY;`,
	"IFCSANITARYTERMINAL": `ENTITY IfcSanitaryTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcSanitaryTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSanitaryTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSanitaryTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSANITARYTERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSANITARYTERMINALTYPE": `ENTITY IfcSanitaryTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcSanitaryTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSanitaryTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSanitaryTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSCHEDULINGTIME": `ENTITY IfcSchedulingTime
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcEventTime
	,IfcLagTime
	,IfcResourceTime
	,IfcTaskTime
	,IfcWorkTime));
	Name : OPTIONAL IfcLabel;
	DataOrigin : OPTIONAL IfcDataOriginEnum;
	UserDefinedDataOrigin : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCSEAMCURVE": `ENTITY IfcSeamCurve
 SUBTYPE OF (IfcSurfaceCurve);
 WHERE
	SameSurface : IfcAssociatedSurface(SELF\IfcSurfaceCurve.AssociatedGeometry[1]) = IfcAssociatedSurface(SELF\IfcSurfaceCurve.AssociatedGeometry[2]);
	TwoPCurves : SIZEOF(SELF\IfcSurfaceCurve.AssociatedGeometry) = 2;
END_ENTITY;`,
	"IFCSECONDORDERPOLYNOMIALSPIRAL": `ENTITY IfcSecondOrderPolynomialSpiral
 SUBTYPE OF (IfcSpiral);
	QuadraticTerm : IfcLengthMeasure;
	LinearTerm : OPTIONAL IfcLengthMeasure;
	ConstantTerm : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCSECTIONEDSOLID": `ENTITY IfcSectionedSolid
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcSectionedSolidHorizontal))
 SUBTYPE OF (IfcSolidModel);
	Directrix : IfcCurve;
	CrossSections : LIST [2:?] OF IfcProfileDef;
 WHERE
	ConsistentProfileTypes : SIZEOF(QUERY(temp <* CrossSections | CrossSections[1].ProfileType <> temp.ProfileType)) = 0;
	DirectrixIs3D : Directrix.Dim = 3;
	SectionsSameType : SIZEOF(QUERY(temp <* CrossSections | TYPEOF(CrossSections[1]) :<>: TYPEOF(temp))) = 0;
END_ENTITY;`,
	"IFCSECTIONEDSOLIDHORIZONTAL": `ENTITY IfcSectionedSolidHorizontal
 SUBTYPE OF (IfcSectionedSolid);
	CrossSectionPositions : LIST [2:?] OF IfcAxis2PlacementLinear;
 WHERE
	CorrespondingSectionPositions : SIZEOF(CrossSections) = SIZEOF(CrossSectionPositions);
	NoLongitudinalOffsets : SIZEOF(QUERY(temp <* CrossSectionPositions | EXISTS(temp.Location.OffsetLongitudinal))) = 0;
END_ENTITY;`,
	"IFCSECTIONEDSPINE": `ENTITY IfcSectionedSpine
 SUBTYPE OF (IfcGeometricRepresentationItem);
	SpineCurve : IfcCompositeCurve;
	CrossSections : LIST [2:?] OF IfcProfileDef;
	CrossSectionPositions : LIST [2:?] OF IfcAxis2Placement3D;
 DERIVE
	 Dim : IfcDimensionCount := 3;
 WHERE
	ConsistentProfileTypes : SIZEOF(QUERY(temp <* CrossSections | CrossSections[1].ProfileType <> temp.ProfileType)) = 0;
	CorrespondingSectionPositions : SIZEOF(CrossSections) = SIZEOF(CrossSectionPositions);
	SpineCurveDim : SpineCurve.Dim = 3;
END_ENTITY;`,
	"IFCSECTIONEDSURFACE": `ENTITY IfcSectionedSurface
 SUBTYPE OF (IfcSurface);
	Directrix : IfcCurve;
	CrossSectionPositions : LIST [2:?] OF IfcAxis2PlacementLinear;
	CrossSections : LIST [2:?] OF IfcProfileDef;
 WHERE
	AreaProfileTypes : SIZEOF(QUERY(temp <* CrossSections | temp.ProfileType = IfcProfileTypeEnum.CURVE)) <> 0;
	CorrespondingSectionPositions : SIZEOF(CrossSections) = SIZEOF(CrossSectionPositions);
	DirectrixIs3D : Directrix.Dim = 3;
	NoOffsets : SIZEOF(QUERY(temp <* CrossSectionPositions | EXISTS(temp.Location.OffsetLateral) OR EXISTS(temp.Location.OffsetVertical) OR EXISTS(temp.Location.OffsetLongitudinal))) = 0;
	SectionsSameType : SIZEOF(QUERY(temp <* CrossSections | TYPEOF(CrossSections[1]) :<>: TYPEOF(temp))) = 0;
END_ENTITY;`,
	"IFCSECTIONPROPERTIES": `ENTITY IfcSectionProperties
 SUBTYPE OF (IfcPreDefinedProperties);
	SectionType : IfcSectionTypeEnum;
	StartProfile : IfcProfileDef;
	EndProfile : OPTIONAL IfcProfileDef;
END_ENTITY;`,
	"IFCSECTIONREINFORCEMENTPROPERTIES": `ENTITY IfcSectionReinforcementProperties
 SUBTYPE OF (IfcPreDefinedProperties);
	LongitudinalStartPosition : IfcLengthMeasure;
	LongitudinalEndPosition : IfcLengthMeasure;
	TransversePosition : OPTIONAL IfcLengthMeasure;
	ReinforcementRole : IfcReinforcingBarRoleEnum;
	SectionDefinition : IfcSectionProperties;
	CrossSectionReinforcementDefinitions : SET [1:?] OF IfcReinforcementBarProperties;
END_ENTITY;`,
	"IFCSEGMENT": `ENTITY IfcSegment
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCompositeCurveSegment
	,IfcCurveSegment))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Transition : IfcTransitionCode;
 DERIVE
	Dim : IfcDimensionCount := IfcSegmentDim(SELF);
 INVERSE
	UsingCurves : SET [1:?] OF IfcCompositeCurve FOR Segments;
END_ENTITY;`,
	"IFCSEGMENTEDREFERENCECURVE": `ENTITY IfcSegmentedReferenceCurve
 SUBTYPE OF (IfcCompositeCurve);
	BaseCurve : IfcBoundedCurve;
	EndPoint : OPTIONAL IfcPlacement;
END_ENTITY;`,
	"IFCSENSOR": `ENTITY IfcSensor
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcSensorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSensorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSensorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSENSORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSENSORTYPE": `ENTITY IfcSensorType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcSensorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSensorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSensorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSEVENTHORDERPOLYNOMIALSPIRAL": `ENTITY IfcSeventhOrderPolynomialSpiral
 SUBTYPE OF (IfcSpiral);
	SepticTerm : IfcLengthMeasure;
	SexticTerm : OPTIONAL IfcLengthMeasure;
	QuinticTerm : OPTIONAL IfcLengthMeasure;
	QuarticTerm : OPTIONAL IfcLengthMeasure;
	CubicTerm : OPTIONAL IfcLengthMeasure;
	QuadraticTerm : OPTIONAL IfcLengthMeasure;
	LinearTerm : OPTIONAL IfcLengthMeasure;
	ConstantTerm : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCSHADINGDEVICE": `ENTITY IfcShadingDevice
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcShadingDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcShadingDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcShadingDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSHADINGDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSHADINGDEVICETYPE": `ENTITY IfcShadingDeviceType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcShadingDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcShadingDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcShadingDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSHAPEASPECT": `ENTITY IfcShapeAspect;
	ShapeRepresentations : LIST [1:?] OF IfcShapeModel;
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	ProductDefinitional : IfcLogical;
	PartOfProductDefinitionShape : OPTIONAL IfcProductRepresentationSelect;
 INVERSE
	HasExternalReferences : SET [0:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCSHAPEMODEL": `ENTITY IfcShapeModel
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcShapeRepresentation
	,IfcTopologyRepresentation))
 SUBTYPE OF (IfcRepresentation);
 INVERSE
	OfShapeAspect : SET [0:1] OF IfcShapeAspect FOR ShapeRepresentations;
 WHERE
	WR11 : (SIZEOF(SELF\IfcRepresentation.OfProductRepresentation) = 1) XOR 
(SIZEOF(SELF\IfcRepresentation.RepresentationMap) = 1) XOR
(SIZEOF(OfShapeAspect) = 1);
END_ENTITY;`,
	"IFCSHAPEREPRESENTATION": `ENTITY IfcShapeRepresentation
 SUBTYPE OF (IfcShapeModel);
 WHERE
	CorrectContext : 'IFC4X3_DEV_73740fe4.IFCGEOMETRICREPRESENTATIONCONTEXT' 
IN TYPEOF(SELF\IfcRepresentation.ContextOfItems);
	CorrectItemsForType : IfcShapeRepresentationTypes(SELF\IfcRepresentation.RepresentationType, SELF\IfcRepresentation.Items);
	HasRepresentationIdentifier : EXISTS(SELF\IfcRepresentation.RepresentationIdentifier);
	HasRepresentationType : EXISTS(SELF\IfcRepresentation.RepresentationType);
	NoTopologicalItem : SIZEOF(QUERY(temp <* Items | 
  ('IFC4X3_DEV_73740fe4.IFCTOPOLOGICALREPRESENTATIONITEM' IN TYPEOF(temp))
    AND (NOT(SIZEOF(
      ['IFC4X3_DEV_73740fe4.IFCVERTEXPOINT',
       'IFC4X3_DEV_73740fe4.IFCEDGECURVE',
       'IFC4X3_DEV_73740fe4.IFCFACESURFACE'] * TYPEOF(temp)) = 1))
)) = 0;
END_ENTITY;`,
	"IFCSHELLBASEDSURFACEMODEL": `ENTITY IfcShellBasedSurfaceModel
 SUBTYPE OF (IfcGeometricRepresentationItem);
	SbsmBoundary : SET [1:?] OF IfcShell;
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCSIGN": `ENTITY IfcSign
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcSignTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSignTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSignTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSIGNTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSIGNAL": `ENTITY IfcSignal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcSignalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSignalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSignalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSIGNALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSIGNALTYPE": `ENTITY IfcSignalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcSignalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSignalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSignalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSIGNTYPE": `ENTITY IfcSignType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcSignTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSignTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSignTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSIMPLEPROPERTY": `ENTITY IfcSimpleProperty
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPropertyBoundedValue
	,IfcPropertyEnumeratedValue
	,IfcPropertyListValue
	,IfcPropertyReferenceValue
	,IfcPropertySingleValue
	,IfcPropertyTableValue))
 SUBTYPE OF (IfcProperty);
END_ENTITY;`,
	"IFCSIMPLEPROPERTYTEMPLATE": `ENTITY IfcSimplePropertyTemplate
 SUBTYPE OF (IfcPropertyTemplate);
	TemplateType : OPTIONAL IfcSimplePropertyTemplateTypeEnum;
	PrimaryMeasureType : OPTIONAL IfcLabel;
	SecondaryMeasureType : OPTIONAL IfcLabel;
	Enumerators : OPTIONAL IfcPropertyEnumeration;
	PrimaryUnit : OPTIONAL IfcUnit;
	SecondaryUnit : OPTIONAL IfcUnit;
	Expression : OPTIONAL IfcLabel;
	AccessState : OPTIONAL IfcStateEnum;
END_ENTITY;`,
	"IFCSINESPIRAL": `ENTITY IfcSineSpiral
 SUBTYPE OF (IfcSpiral);
	SineTerm : IfcLengthMeasure;
	LinearTerm : OPTIONAL IfcLengthMeasure;
	ConstantTerm : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCSITE": `ENTITY IfcSite
 SUBTYPE OF (IfcSpatialStructureElement);
	RefLatitude : OPTIONAL IfcCompoundPlaneAngleMeasure;
	RefLongitude : OPTIONAL IfcCompoundPlaneAngleMeasure;
	RefElevation : OPTIONAL IfcLengthMeasure;
	LandTitleNumber : OPTIONAL IfcLabel;
	SiteAddress : OPTIONAL IfcPostalAddress;
END_ENTITY;`,
	"IFCSIUNIT": `ENTITY IfcSIUnit
 SUBTYPE OF (IfcNamedUnit);
	Prefix : OPTIONAL IfcSIPrefix;
	Name : IfcSIUnitName;
 DERIVE
	 SELF\IfcNamedUnit.Dimensions : IfcDimensionalExponents := IfcDimensionsForSIUnit (SELF.Name);
END_ENTITY;`,
	"IFCSLAB": `ENTITY IfcSlab
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcSlabTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSlabTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSlabTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSLABTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSLABTYPE": `ENTITY IfcSlabType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcSlabTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSlabTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSlabTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSLIPPAGECONNECTIONCONDITION": `ENTITY IfcSlippageConnectionCondition
 SUBTYPE OF (IfcStructuralConnectionCondition);
	SlippageX : OPTIONAL IfcLengthMeasure;
	SlippageY : OPTIONAL IfcLengthMeasure;
	SlippageZ : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCSOLARDEVICE": `ENTITY IfcSolarDevice
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcSolarDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSolarDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSolarDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSOLARDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSOLARDEVICETYPE": `ENTITY IfcSolarDeviceType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcSolarDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSolarDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSolarDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSOLIDMODEL": `ENTITY IfcSolidModel
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcCsgSolid
	,IfcManifoldSolidBrep
	,IfcSectionedSolid
	,IfcSweptAreaSolid
	,IfcSweptDiskSolid))
 SUBTYPE OF (IfcGeometricRepresentationItem);
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCSPACE": `ENTITY IfcSpace
 SUBTYPE OF (IfcSpatialStructureElement);
	PredefinedType : OPTIONAL IfcSpaceTypeEnum;
	ElevationWithFlooring : OPTIONAL IfcLengthMeasure;
 INVERSE
	HasCoverings : SET [0:?] OF IfcRelCoversSpaces FOR RelatingSpace;
	BoundedBy : SET [0:?] OF IfcRelSpaceBoundary FOR RelatingSpace;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSpaceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSpaceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSPACETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSPACEHEATER": `ENTITY IfcSpaceHeater
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcSpaceHeaterTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSpaceHeaterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSpaceHeaterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSPACEHEATERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSPACEHEATERTYPE": `ENTITY IfcSpaceHeaterType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcSpaceHeaterTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSpaceHeaterTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSpaceHeaterTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSPACETYPE": `ENTITY IfcSpaceType
 SUBTYPE OF (IfcSpatialStructureElementType);
	PredefinedType : IfcSpaceTypeEnum;
	LongName : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSpaceTypeEnum.USERDEFINED) OR 
((PredefinedType = IfcSpaceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcSpatialElementType.ElementType));
END_ENTITY;`,
	"IFCSPATIALELEMENT": `ENTITY IfcSpatialElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcExternalSpatialStructureElement
	,IfcSpatialStructureElement
	,IfcSpatialZone))
 SUBTYPE OF (IfcProduct);
	LongName : OPTIONAL IfcLabel;
 INVERSE
	ContainsElements : SET [0:?] OF IfcRelContainedInSpatialStructure FOR RelatingStructure;
	ServicedBySystems : SET [0:?] OF IfcRelServicesBuildings FOR RelatedBuildings;
	ReferencesElements : SET [0:?] OF IfcRelReferencedInSpatialStructure FOR RelatingStructure;
	IsInterferedByElements : SET [0:?] OF IfcRelInterferesElements FOR RelatedElement;
	InterferesElements : SET [0:?] OF IfcRelInterferesElements FOR RelatingElement;
END_ENTITY;`,
	"IFCSPATIALELEMENTTYPE": `ENTITY IfcSpatialElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcSpatialStructureElementType
	,IfcSpatialZoneType))
 SUBTYPE OF (IfcTypeProduct);
	ElementType : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCSPATIALSTRUCTUREELEMENT": `ENTITY IfcSpatialStructureElement
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBuildingStorey
	,IfcFacility
	,IfcFacilityPart
	,IfcSite
	,IfcSpace))
 SUBTYPE OF (IfcSpatialElement);
	CompositionType : OPTIONAL IfcElementCompositionEnum;
 WHERE
	WR41 : (HIINDEX(SELF\IfcObjectDefinition.Decomposes) = 1) 
AND
('IFC4X3_DEV_73740fe4.IFCRELAGGREGATES' IN TYPEOF(SELF\IfcObjectDefinition.Decomposes[1])) 
AND
(('IFC4X3_DEV_73740fe4.IFCPROJECT' IN TYPEOF (SELF\IfcObjectDefinition.Decomposes[1].RelatingObject)) OR
 ('IFC4X3_DEV_73740fe4.IFCSPATIALSTRUCTUREELEMENT' IN TYPEOF (SELF\IfcObjectDefinition.Decomposes[1].RelatingObject))
);
END_ENTITY;`,
	"IFCSPATIALSTRUCTUREELEMENTTYPE": `ENTITY IfcSpatialStructureElementType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcSpaceType))
 SUBTYPE OF (IfcSpatialElementType);
END_ENTITY;`,
	"IFCSPATIALZONE": `ENTITY IfcSpatialZone
 SUBTYPE OF (IfcSpatialElement);
	PredefinedType : OPTIONAL IfcSpatialZoneTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSpatialZoneTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSpatialZoneTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSPATIALZONETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSPATIALZONETYPE": `ENTITY IfcSpatialZoneType
 SUBTYPE OF (IfcSpatialElementType);
	PredefinedType : IfcSpatialZoneTypeEnum;
	LongName : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSpatialZoneTypeEnum.USERDEFINED) OR 
((PredefinedType = IfcSpatialZoneTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcSpatialElementType.ElementType));
END_ENTITY;`,
	"IFCSPHERE": `ENTITY IfcSphere
 SUBTYPE OF (IfcCsgPrimitive3D);
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCSPHERICALSURFACE": `ENTITY IfcSphericalSurface
 SUBTYPE OF (IfcElementarySurface);
	Radius : IfcPositiveLengthMeasure;
END_ENTITY;`,
	"IFCSPIRAL": `ENTITY IfcSpiral
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcClothoid
	,IfcCosineSpiral
	,IfcSecondOrderPolynomialSpiral
	,IfcSeventhOrderPolynomialSpiral
	,IfcSineSpiral
	,IfcThirdOrderPolynomialSpiral))
 SUBTYPE OF (IfcCurve);
	Position : IfcAxis2Placement;
END_ENTITY;`,
	"IFCSTACKTERMINAL": `ENTITY IfcStackTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcStackTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcStackTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStackTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSTACKTERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSTACKTERMINALTYPE": `ENTITY IfcStackTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcStackTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcStackTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStackTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSTAIR": `ENTITY IfcStair
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcStairTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcStairTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStairTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSTAIRTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSTAIRFLIGHT": `ENTITY IfcStairFlight
 SUBTYPE OF (IfcBuiltElement);
	NumberOfRisers : OPTIONAL IfcInteger;
	NumberOfTreads : OPTIONAL IfcInteger;
	RiserHeight : OPTIONAL IfcPositiveLengthMeasure;
	TreadLength : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcStairFlightTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcStairFlightTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStairFlightTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSTAIRFLIGHTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSTAIRFLIGHTTYPE": `ENTITY IfcStairFlightType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcStairFlightTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcStairFlightTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStairFlightTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSTAIRTYPE": `ENTITY IfcStairType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcStairTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcStairTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcStairTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSTRUCTURALACTION": `ENTITY IfcStructuralAction
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralCurveAction
	,IfcStructuralPointAction
	,IfcStructuralSurfaceAction))
 SUBTYPE OF (IfcStructuralActivity);
	DestabilizingLoad : OPTIONAL IfcBoolean;
END_ENTITY;`,
	"IFCSTRUCTURALACTIVITY": `ENTITY IfcStructuralActivity
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralAction
	,IfcStructuralReaction))
 SUBTYPE OF (IfcProduct);
	AppliedLoad : IfcStructuralLoad;
	GlobalOrLocal : IfcGlobalOrLocalEnum;
 INVERSE
	AssignedToStructuralItem : SET [0:1] OF IfcRelConnectsStructuralActivity FOR RelatedStructuralActivity;
END_ENTITY;`,
	"IFCSTRUCTURALANALYSISMODEL": `ENTITY IfcStructuralAnalysisModel
 SUBTYPE OF (IfcSystem);
	PredefinedType : IfcAnalysisModelTypeEnum;
	OrientationOf2DPlane : OPTIONAL IfcAxis2Placement3D;
	LoadedBy : OPTIONAL SET [1:?] OF IfcStructuralLoadGroup;
	HasResults : OPTIONAL SET [1:?] OF IfcStructuralResultGroup;
	SharedPlacement : OPTIONAL IfcObjectPlacement;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcAnalysisModelTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcAnalysisModelTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCSTRUCTURALCONNECTION": `ENTITY IfcStructuralConnection
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralCurveConnection
	,IfcStructuralPointConnection
	,IfcStructuralSurfaceConnection))
 SUBTYPE OF (IfcStructuralItem);
	AppliedCondition : OPTIONAL IfcBoundaryCondition;
 INVERSE
	ConnectsStructuralMembers : SET [1:?] OF IfcRelConnectsStructuralMember FOR RelatedStructuralConnection;
END_ENTITY;`,
	"IFCSTRUCTURALCONNECTIONCONDITION": `ENTITY IfcStructuralConnectionCondition
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcFailureConnectionCondition
	,IfcSlippageConnectionCondition));
	Name : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCSTRUCTURALCURVEACTION": `ENTITY IfcStructuralCurveAction
 SUPERTYPE OF (ONEOF
	(IfcStructuralLinearAction))
 SUBTYPE OF (IfcStructuralAction);
	ProjectedOrTrue : OPTIONAL IfcProjectedOrTrueLengthEnum;
	PredefinedType : IfcStructuralCurveActivityTypeEnum;
 WHERE
	HasObjectType : (PredefinedType <> IfcStructuralCurveActivityTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
	ProjectedIsGlobal : (NOT EXISTS(ProjectedOrTrue)) OR
((ProjectedOrTrue <> PROJECTED_LENGTH) OR
 (SELF\IfcStructuralActivity.GlobalOrLocal = GLOBAL_COORDS));
	SuitablePredefinedType : PredefinedType <> IfcStructuralCurveActivityTypeEnum.EQUIDISTANT;
END_ENTITY;`,
	"IFCSTRUCTURALCURVECONNECTION": `ENTITY IfcStructuralCurveConnection
 SUBTYPE OF (IfcStructuralConnection);
	AxisDirection : IfcDirection;
END_ENTITY;`,
	"IFCSTRUCTURALCURVEMEMBER": `ENTITY IfcStructuralCurveMember
 SUPERTYPE OF (ONEOF
	(IfcStructuralCurveMemberVarying))
 SUBTYPE OF (IfcStructuralMember);
	PredefinedType : IfcStructuralCurveMemberTypeEnum;
	Axis : IfcDirection;
 WHERE
	HasObjectType : (PredefinedType <> IfcStructuralCurveMemberTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCSTRUCTURALCURVEMEMBERVARYING": `ENTITY IfcStructuralCurveMemberVarying
 SUBTYPE OF (IfcStructuralCurveMember);
END_ENTITY;`,
	"IFCSTRUCTURALCURVEREACTION": `ENTITY IfcStructuralCurveReaction
 SUBTYPE OF (IfcStructuralReaction);
	PredefinedType : IfcStructuralCurveActivityTypeEnum;
 WHERE
	HasObjectType : (PredefinedType <> IfcStructuralCurveActivityTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
	SuitablePredefinedType : (PredefinedType <> IfcStructuralCurveActivityTypeEnum.SINUS) AND (PredefinedType <> IfcStructuralCurveActivityTypeEnum.PARABOLA);
END_ENTITY;`,
	"IFCSTRUCTURALITEM": `ENTITY IfcStructuralItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralConnection
	,IfcStructuralMember))
 SUBTYPE OF (IfcProduct);
 INVERSE
	AssignedStructuralActivity : SET [0:?] OF IfcRelConnectsStructuralActivity FOR RelatingElement;
END_ENTITY;`,
	"IFCSTRUCTURALLINEARACTION": `ENTITY IfcStructuralLinearAction
 SUBTYPE OF (IfcStructuralCurveAction);
 WHERE
	ConstPredefinedType : SELF\IfcStructuralCurveAction.PredefinedType = IfcStructuralCurveActivityTypeEnum.CONST;
	SuitableLoadType : SIZEOF(['IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADLINEARFORCE', 'IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADTEMPERATURE'] * TYPEOF(SELF\IfcStructuralActivity.AppliedLoad)) = 1;
END_ENTITY;`,
	"IFCSTRUCTURALLOAD": `ENTITY IfcStructuralLoad
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralLoadConfiguration
	,IfcStructuralLoadOrResult));
	Name : OPTIONAL IfcLabel;
END_ENTITY;`,
	"IFCSTRUCTURALLOADCASE": `ENTITY IfcStructuralLoadCase
 SUBTYPE OF (IfcStructuralLoadGroup);
	SelfWeightCoefficients : OPTIONAL LIST [3:3] OF IfcRatioMeasure;
 WHERE
	IsLoadCasePredefinedType : SELF\IfcStructuralLoadGroup.PredefinedType = IfcLoadGroupTypeEnum.LOAD_CASE;
END_ENTITY;`,
	"IFCSTRUCTURALLOADCONFIGURATION": `ENTITY IfcStructuralLoadConfiguration
 SUBTYPE OF (IfcStructuralLoad);
	Values : LIST [1:?] OF IfcStructuralLoadOrResult;
	Locations : OPTIONAL LIST [1:?] OF UNIQUE LIST [1:2] OF IfcLengthMeasure;
 WHERE
	ValidListSize : NOT EXISTS(Locations) OR (SIZEOF(Locations) = SIZEOF(Values));
END_ENTITY;`,
	"IFCSTRUCTURALLOADGROUP": `ENTITY IfcStructuralLoadGroup
 SUPERTYPE OF (ONEOF
	(IfcStructuralLoadCase))
 SUBTYPE OF (IfcGroup);
	PredefinedType : IfcLoadGroupTypeEnum;
	ActionType : IfcActionTypeEnum;
	ActionSource : IfcActionSourceTypeEnum;
	Coefficient : OPTIONAL IfcRatioMeasure;
	Purpose : OPTIONAL IfcLabel;
 INVERSE
	SourceOfResultGroup : SET [0:1] OF IfcStructuralResultGroup FOR ResultForLoadGroup;
	LoadGroupFor : SET [0:?] OF IfcStructuralAnalysisModel FOR LoadedBy;
 WHERE
	HasObjectType : (
  (PredefinedType <> IfcLoadGroupTypeEnum.USERDEFINED) AND
  (ActionType <> IfcActionTypeEnum.USERDEFINED) AND
  (ActionSource <> IfcActionSourceTypeEnum.USERDEFINED)
) OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCSTRUCTURALLOADLINEARFORCE": `ENTITY IfcStructuralLoadLinearForce
 SUBTYPE OF (IfcStructuralLoadStatic);
	LinearForceX : OPTIONAL IfcLinearForceMeasure;
	LinearForceY : OPTIONAL IfcLinearForceMeasure;
	LinearForceZ : OPTIONAL IfcLinearForceMeasure;
	LinearMomentX : OPTIONAL IfcLinearMomentMeasure;
	LinearMomentY : OPTIONAL IfcLinearMomentMeasure;
	LinearMomentZ : OPTIONAL IfcLinearMomentMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADORRESULT": `ENTITY IfcStructuralLoadOrResult
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralLoadStatic
	,IfcSurfaceReinforcementArea))
 SUBTYPE OF (IfcStructuralLoad);
END_ENTITY;`,
	"IFCSTRUCTURALLOADPLANARFORCE": `ENTITY IfcStructuralLoadPlanarForce
 SUBTYPE OF (IfcStructuralLoadStatic);
	PlanarForceX : OPTIONAL IfcPlanarForceMeasure;
	PlanarForceY : OPTIONAL IfcPlanarForceMeasure;
	PlanarForceZ : OPTIONAL IfcPlanarForceMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADSINGLEDISPLACEMENT": `ENTITY IfcStructuralLoadSingleDisplacement
 SUPERTYPE OF (ONEOF
	(IfcStructuralLoadSingleDisplacementDistortion))
 SUBTYPE OF (IfcStructuralLoadStatic);
	DisplacementX : OPTIONAL IfcLengthMeasure;
	DisplacementY : OPTIONAL IfcLengthMeasure;
	DisplacementZ : OPTIONAL IfcLengthMeasure;
	RotationalDisplacementRX : OPTIONAL IfcPlaneAngleMeasure;
	RotationalDisplacementRY : OPTIONAL IfcPlaneAngleMeasure;
	RotationalDisplacementRZ : OPTIONAL IfcPlaneAngleMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADSINGLEDISPLACEMENTDISTORTION": `ENTITY IfcStructuralLoadSingleDisplacementDistortion
 SUBTYPE OF (IfcStructuralLoadSingleDisplacement);
	Distortion : OPTIONAL IfcCurvatureMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADSINGLEFORCE": `ENTITY IfcStructuralLoadSingleForce
 SUPERTYPE OF (ONEOF
	(IfcStructuralLoadSingleForceWarping))
 SUBTYPE OF (IfcStructuralLoadStatic);
	ForceX : OPTIONAL IfcForceMeasure;
	ForceY : OPTIONAL IfcForceMeasure;
	ForceZ : OPTIONAL IfcForceMeasure;
	MomentX : OPTIONAL IfcTorqueMeasure;
	MomentY : OPTIONAL IfcTorqueMeasure;
	MomentZ : OPTIONAL IfcTorqueMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADSINGLEFORCEWARPING": `ENTITY IfcStructuralLoadSingleForceWarping
 SUBTYPE OF (IfcStructuralLoadSingleForce);
	WarpingMoment : OPTIONAL IfcWarpingMomentMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALLOADSTATIC": `ENTITY IfcStructuralLoadStatic
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralLoadLinearForce
	,IfcStructuralLoadPlanarForce
	,IfcStructuralLoadSingleDisplacement
	,IfcStructuralLoadSingleForce
	,IfcStructuralLoadTemperature))
 SUBTYPE OF (IfcStructuralLoadOrResult);
END_ENTITY;`,
	"IFCSTRUCTURALLOADTEMPERATURE": `ENTITY IfcStructuralLoadTemperature
 SUBTYPE OF (IfcStructuralLoadStatic);
	DeltaTConstant : OPTIONAL IfcThermodynamicTemperatureMeasure;
	DeltaTY : OPTIONAL IfcThermodynamicTemperatureMeasure;
	DeltaTZ : OPTIONAL IfcThermodynamicTemperatureMeasure;
END_ENTITY;`,
	"IFCSTRUCTURALMEMBER": `ENTITY IfcStructuralMember
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralCurveMember
	,IfcStructuralSurfaceMember))
 SUBTYPE OF (IfcStructuralItem);
 INVERSE
	ConnectedBy : SET [0:?] OF IfcRelConnectsStructuralMember FOR RelatingStructuralMember;
END_ENTITY;`,
	"IFCSTRUCTURALPLANARACTION": `ENTITY IfcStructuralPlanarAction
 SUBTYPE OF (IfcStructuralSurfaceAction);
 WHERE
	ConstPredefinedType : SELF\IfcStructuralSurfaceAction.PredefinedType = IfcStructuralSurfaceActivityTypeEnum.CONST;
	SuitableLoadType : SIZEOF(['IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADPLANARFORCE', 'IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADTEMPERATURE'] * TYPEOF(SELF\IfcStructuralActivity.AppliedLoad)) = 1;
END_ENTITY;`,
	"IFCSTRUCTURALPOINTACTION": `ENTITY IfcStructuralPointAction
 SUBTYPE OF (IfcStructuralAction);
 WHERE
	SuitableLoadType : SIZEOF(['IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADSINGLEFORCE', 'IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADSINGLEDISPLACEMENT'] * TYPEOF(SELF\IfcStructuralActivity.AppliedLoad)) = 1;
END_ENTITY;`,
	"IFCSTRUCTURALPOINTCONNECTION": `ENTITY IfcStructuralPointConnection
 SUBTYPE OF (IfcStructuralConnection);
	ConditionCoordinateSystem : OPTIONAL IfcAxis2Placement3D;
END_ENTITY;`,
	"IFCSTRUCTURALPOINTREACTION": `ENTITY IfcStructuralPointReaction
 SUBTYPE OF (IfcStructuralReaction);
 WHERE
	SuitableLoadType : SIZEOF(['IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADSINGLEFORCE', 'IFC4X3_DEV_73740fe4.IFCSTRUCTURALLOADSINGLEDISPLACEMENT'] * TYPEOF(SELF\IfcStructuralActivity.AppliedLoad)) = 1;
END_ENTITY;`,
	"IFCSTRUCTURALREACTION": `ENTITY IfcStructuralReaction
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStructuralCurveReaction
	,IfcStructuralPointReaction
	,IfcStructuralSurfaceReaction))
 SUBTYPE OF (IfcStructuralActivity);
END_ENTITY;`,
	"IFCSTRUCTURALRESULTGROUP": `ENTITY IfcStructuralResultGroup
 SUBTYPE OF (IfcGroup);
	TheoryType : IfcAnalysisTheoryTypeEnum;
	ResultForLoadGroup : OPTIONAL IfcStructuralLoadGroup;
	IsLinear : IfcBoolean;
 INVERSE
	ResultGroupFor : SET [0:1] OF IfcStructuralAnalysisModel FOR HasResults;
 WHERE
	HasObjectType : (TheoryType <> IfcAnalysisTheoryTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCSTRUCTURALSURFACEACTION": `ENTITY IfcStructuralSurfaceAction
 SUPERTYPE OF (ONEOF
	(IfcStructuralPlanarAction))
 SUBTYPE OF (IfcStructuralAction);
	ProjectedOrTrue : OPTIONAL IfcProjectedOrTrueLengthEnum;
	PredefinedType : IfcStructuralSurfaceActivityTypeEnum;
 WHERE
	HasObjectType : (PredefinedType <> IfcStructuralSurfaceActivityTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
	ProjectedIsGlobal : (NOT EXISTS(ProjectedOrTrue)) OR
((ProjectedOrTrue <> PROJECTED_LENGTH) OR
 (SELF\IfcStructuralActivity.GlobalOrLocal = GLOBAL_COORDS));
END_ENTITY;`,
	"IFCSTRUCTURALSURFACECONNECTION": `ENTITY IfcStructuralSurfaceConnection
 SUBTYPE OF (IfcStructuralConnection);
END_ENTITY;`,
	"IFCSTRUCTURALSURFACEMEMBER": `ENTITY IfcStructuralSurfaceMember
 SUPERTYPE OF (ONEOF
	(IfcStructuralSurfaceMemberVarying))
 SUBTYPE OF (IfcStructuralMember);
	PredefinedType : IfcStructuralSurfaceMemberTypeEnum;
	Thickness : OPTIONAL IfcPositiveLengthMeasure;
 WHERE
	HasObjectType : (PredefinedType <> IfcStructuralSurfaceMemberTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCSTRUCTURALSURFACEMEMBERVARYING": `ENTITY IfcStructuralSurfaceMemberVarying
 SUBTYPE OF (IfcStructuralSurfaceMember);
END_ENTITY;`,
	"IFCSTRUCTURALSURFACEREACTION": `ENTITY IfcStructuralSurfaceReaction
 SUBTYPE OF (IfcStructuralReaction);
	PredefinedType : IfcStructuralSurfaceActivityTypeEnum;
 WHERE
	HasPredefinedType : (PredefinedType <> IfcStructuralSurfaceActivityTypeEnum.USERDEFINED) OR EXISTS(SELF\IfcObject.ObjectType);
END_ENTITY;`,
	"IFCSTYLEDITEM": `ENTITY IfcStyledItem
 SUBTYPE OF (IfcRepresentationItem);
	Item : OPTIONAL IfcRepresentationItem;
	Styles : SET [1:?] OF IfcPresentationStyle;
	Name : OPTIONAL IfcLabel;
 WHERE
	ApplicableItem : NOT('IFC4X3_DEV_73740fe4.IFCSTYLEDITEM' IN TYPEOF(Item));
END_ENTITY;`,
	"IFCSTYLEDREPRESENTATION": `ENTITY IfcStyledRepresentation
 SUBTYPE OF (IfcStyleModel);
 WHERE
	OnlyStyledItems : SIZEOF(QUERY(temp <* SELF\IfcRepresentation.Items | 
  (NOT('IFC4X3_DEV_73740fe4.IFCSTYLEDITEM' IN TYPEOF(temp)))
)) = 0;
END_ENTITY;`,
	"IFCSTYLEMODEL": `ENTITY IfcStyleModel
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcStyledRepresentation))
 SUBTYPE OF (IfcRepresentation);
END_ENTITY;`,
	"IFCSUBCONTRACTRESOURCE": `ENTITY IfcSubContractResource
 SUBTYPE OF (IfcConstructionResource);
	PredefinedType : OPTIONAL IfcSubContractResourceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR 
 (PredefinedType <> IfcSubContractResourceTypeEnum.USERDEFINED) OR 
 ((PredefinedType = IfcSubContractResourceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCSUBCONTRACTRESOURCETYPE": `ENTITY IfcSubContractResourceType
 SUBTYPE OF (IfcConstructionResourceType);
	PredefinedType : IfcSubContractResourceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSubContractResourceTypeEnum.USERDEFINED) OR
((PredefinedType = IfcSubContractResourceTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeResource.ResourceType));
END_ENTITY;`,
	"IFCSUBEDGE": `ENTITY IfcSubedge
 SUBTYPE OF (IfcEdge);
	ParentEdge : IfcEdge;
END_ENTITY;`,
	"IFCSURFACE": `ENTITY IfcSurface
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBoundedSurface
	,IfcElementarySurface
	,IfcSectionedSurface
	,IfcSweptSurface))
 SUBTYPE OF (IfcGeometricRepresentationItem);
 DERIVE
	 Dim : IfcDimensionCount := 3;
END_ENTITY;`,
	"IFCSURFACECURVE": `ENTITY IfcSurfaceCurve
 SUPERTYPE OF (ONEOF
	(IfcIntersectionCurve
	,IfcSeamCurve))
 SUBTYPE OF (IfcCurve);
	Curve3D : IfcCurve;
	AssociatedGeometry : LIST [1:2] OF IfcPcurve;
	MasterRepresentation : IfcPreferredSurfaceCurveRepresentation;
 DERIVE
	 BasisSurface : SET [1:2] OF IfcSurface := IfcGetBasisSurface(SELF);
 WHERE
	CurveIs3D : Curve3D.Dim = 3;
	CurveIsNotPcurve : NOT ('IFC4X3_DEV_73740fe4.IFCPCURVE' IN TYPEOF(Curve3D));
END_ENTITY;`,
	"IFCSURFACECURVESWEPTAREASOLID": `ENTITY IfcSurfaceCurveSweptAreaSolid
 SUBTYPE OF (IfcDirectrixCurveSweptAreaSolid);
	ReferenceSurface : IfcSurface;
END_ENTITY;`,
	"IFCSURFACEFEATURE": `ENTITY IfcSurfaceFeature
 SUBTYPE OF (IfcFeatureElement);
	PredefinedType : OPTIONAL IfcSurfaceFeatureTypeEnum;
 INVERSE
	AdheresToElement : IfcRelAdheresToElement FOR RelatedSurfaceFeatures;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSurfaceFeatureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSurfaceFeatureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCSURFACEOFLINEAREXTRUSION": `ENTITY IfcSurfaceOfLinearExtrusion
 SUBTYPE OF (IfcSweptSurface);
	ExtrudedDirection : IfcDirection;
	Depth : IfcLengthMeasure;
 DERIVE
	 ExtrusionAxis : IfcVector := IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcVector (ExtrudedDirection, Depth);
 WHERE
	DepthGreaterZero : Depth > 0.;
END_ENTITY;`,
	"IFCSURFACEOFREVOLUTION": `ENTITY IfcSurfaceOfRevolution
 SUBTYPE OF (IfcSweptSurface);
	AxisPosition : IfcAxis1Placement;
 DERIVE
	 AxisLine : IfcLine := IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcCurve() || IfcLine(AxisPosition.Location, 
  IfcRepresentationItem() || IfcGeometricRepresentationItem () || IfcVector(AxisPosition.Z,1.0));
END_ENTITY;`,
	"IFCSURFACEREINFORCEMENTAREA": `ENTITY IfcSurfaceReinforcementArea
 SUBTYPE OF (IfcStructuralLoadOrResult);
	SurfaceReinforcement1 : OPTIONAL LIST [2:3] OF IfcLengthMeasure;
	SurfaceReinforcement2 : OPTIONAL LIST [2:3] OF IfcLengthMeasure;
	ShearReinforcement : OPTIONAL IfcRatioMeasure;
 WHERE
	NonnegativeArea1 : (NOT EXISTS(SurfaceReinforcement1)) OR (
    (SurfaceReinforcement1[1] >= 0.) AND
    (SurfaceReinforcement1[2] >= 0.) AND
    ((SIZEOF(SurfaceReinforcement1) = 1) OR (SurfaceReinforcement1[1] >= 0.))
);
	NonnegativeArea2 : (NOT EXISTS(SurfaceReinforcement2)) OR (
    (SurfaceReinforcement2[1] >= 0.) AND
    (SurfaceReinforcement2[2] >= 0.) AND
    ((SIZEOF(SurfaceReinforcement2) = 1) OR (SurfaceReinforcement2[1] >= 0.))
);
	NonnegativeArea3 : (NOT EXISTS(ShearReinforcement)) OR (ShearReinforcement >= 0.);
	SurfaceAndOrShearAreaSpecified : EXISTS(SurfaceReinforcement1) OR EXISTS(SurfaceReinforcement2) OR EXISTS(ShearReinforcement);
END_ENTITY;`,
	"IFCSURFACESTYLE": `ENTITY IfcSurfaceStyle
 SUBTYPE OF (IfcPresentationStyle);
	Side : IfcSurfaceSide;
	Styles : SET [1:5] OF IfcSurfaceStyleElementSelect;
 WHERE
	MaxOneExtDefined : SIZEOF(QUERY(Style <* SELF.Styles |
  'IFC4X3_DEV_73740fe4.IFCEXTERNALLYDEFINEDSURFACESTYLE' IN
   TYPEOF(Style)
  )) <= 1;
	MaxOneLighting : SIZEOF(QUERY(Style <* SELF.Styles |
  'IFC4X3_DEV_73740fe4.IFCSURFACESTYLELIGHTING' IN
   TYPEOF(Style)
  )) <= 1;
	MaxOneRefraction : SIZEOF(QUERY(Style <* SELF.Styles |
  'IFC4X3_DEV_73740fe4.IFCSURFACESTYLEREFRACTION' IN
   TYPEOF(Style)
  )) <= 1;
	MaxOneShading : SIZEOF(QUERY(Style <* SELF.Styles |
  'IFC4X3_DEV_73740fe4.IFCSURFACESTYLESHADING' IN
   TYPEOF(Style)
  )) <= 1;
	MaxOneTextures : SIZEOF(QUERY(Style <* SELF.Styles |
  'IFC4X3_DEV_73740fe4.IFCSURFACESTYLEWITHTEXTURES' IN
   TYPEOF(Style)
  )) <= 1;
END_ENTITY;`,
	"IFCSURFACESTYLELIGHTING": `ENTITY IfcSurfaceStyleLighting
 SUBTYPE OF (IfcPresentationItem);
	DiffuseTransmissionColour : IfcColourRgb;
	DiffuseReflectionColour : IfcColourRgb;
	TransmissionColour : IfcColourRgb;
	ReflectanceColour : IfcColourRgb;
END_ENTITY;`,
	"IFCSURFACESTYLEREFRACTION": `ENTITY IfcSurfaceStyleRefraction
 SUBTYPE OF (IfcPresentationItem);
	RefractionIndex : OPTIONAL IfcReal;
	DispersionFactor : OPTIONAL IfcReal;
END_ENTITY;`,
	"IFCSURFACESTYLERENDERING": `ENTITY IfcSurfaceStyleRendering
 SUBTYPE OF (IfcSurfaceStyleShading);
	DiffuseColour : OPTIONAL IfcColourOrFactor;
	TransmissionColour : OPTIONAL IfcColourOrFactor;
	DiffuseTransmissionColour : OPTIONAL IfcColourOrFactor;
	ReflectionColour : OPTIONAL IfcColourOrFactor;
	SpecularColour : OPTIONAL IfcColourOrFactor;
	SpecularHighlight : OPTIONAL IfcSpecularHighlightSelect;
	ReflectanceMethod : IfcReflectanceMethodEnum;
END_ENTITY;`,
	"IFCSURFACESTYLESHADING": `ENTITY IfcSurfaceStyleShading
 SUPERTYPE OF (ONEOF
	(IfcSurfaceStyleRendering))
 SUBTYPE OF (IfcPresentationItem);
	SurfaceColour : IfcColourRgb;
	Transparency : OPTIONAL IfcNormalisedRatioMeasure;
END_ENTITY;`,
	"IFCSURFACESTYLEWITHTEXTURES": `ENTITY IfcSurfaceStyleWithTextures
 SUBTYPE OF (IfcPresentationItem);
	Textures : LIST [1:?] OF IfcSurfaceTexture;
END_ENTITY;`,
	"IFCSURFACETEXTURE": `ENTITY IfcSurfaceTexture
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcBlobTexture
	,IfcImageTexture
	,IfcPixelTexture))
 SUBTYPE OF (IfcPresentationItem);
	RepeatS : IfcBoolean;
	RepeatT : IfcBoolean;
	Mode : OPTIONAL IfcIdentifier;
	TextureTransform : OPTIONAL IfcCartesianTransformationOperator2D;
	Parameter : OPTIONAL LIST [1:?] OF IfcIdentifier;
 INVERSE
	IsMappedBy : SET [0:?] OF IfcTextureCoordinate FOR Maps;
	UsedInStyles : SET [0:?] OF IfcSurfaceStyleWithTextures FOR Textures;
END_ENTITY;`,
	"IFCSWEPTAREASOLID": `ENTITY IfcSweptAreaSolid
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcDirectrixCurveSweptAreaSolid
	,IfcExtrudedAreaSolid
	,IfcRevolvedAreaSolid))
 SUBTYPE OF (IfcSolidModel);
	SweptArea : IfcProfileDef;
	Position : OPTIONAL IfcAxis2Placement3D;
 WHERE
	SweptAreaType : SweptArea.ProfileType = IfcProfileTypeEnum.Area;
END_ENTITY;`,
	"IFCSWEPTDISKSOLID": `ENTITY IfcSweptDiskSolid
 SUPERTYPE OF (ONEOF
	(IfcSweptDiskSolidPolygonal))
 SUBTYPE OF (IfcSolidModel);
	Directrix : IfcCurve;
	Radius : IfcPositiveLengthMeasure;
	InnerRadius : OPTIONAL IfcPositiveLengthMeasure;
	StartParam : OPTIONAL IfcParameterValue;
	EndParam : OPTIONAL IfcParameterValue;
 WHERE
	DirectrixBounded : (EXISTS(StartParam) AND EXISTS(EndParam)) OR 
(SIZEOF(['IFC4X3_DEV_73740fe4.IFCCONIC', 'IFC4X3_DEV_73740fe4.IFCBOUNDEDCURVE'] * TYPEOF(Directrix)) = 1);
	DirectrixDim : Directrix.Dim = 3;
	InnerRadiusSize : (NOT EXISTS(InnerRadius)) OR (Radius > InnerRadius);
END_ENTITY;`,
	"IFCSWEPTDISKSOLIDPOLYGONAL": `ENTITY IfcSweptDiskSolidPolygonal
 SUBTYPE OF (IfcSweptDiskSolid);
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
 WHERE
	CorrectRadii : NOT(EXISTS(FilletRadius)) OR (FilletRadius >= SELF\IfcSweptDiskSolid.Radius);
	DirectrixIsPolyline : ('IFC4X3_DEV_73740fe4.IFCPOLYLINE' IN TYPEOF(SELF\IfcSweptDiskSolid.Directrix)) OR
(('IFC4X3_DEV_73740fe4.IFCINDEXEDPOLYCURVE' IN TYPEOF(SELF\IfcSweptDiskSolid.Directrix)) AND NOT(EXISTS(SELF\IfcSweptDiskSolid.Directrix\IfcIndexedPolyCurve.Segments)));
END_ENTITY;`,
	"IFCSWEPTSURFACE": `ENTITY IfcSweptSurface
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcSurfaceOfLinearExtrusion
	,IfcSurfaceOfRevolution))
 SUBTYPE OF (IfcSurface);
	SweptCurve : IfcProfileDef;
	Position : OPTIONAL IfcAxis2Placement3D;
 WHERE
	SweptCurveType : SweptCurve.ProfileType = IfcProfileTypeEnum.Curve;
END_ENTITY;`,
	"IFCSWITCHINGDEVICE": `ENTITY IfcSwitchingDevice
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcSwitchingDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSwitchingDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSwitchingDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSWITCHINGDEVICETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSWITCHINGDEVICETYPE": `ENTITY IfcSwitchingDeviceType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcSwitchingDeviceTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcSwitchingDeviceTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSwitchingDeviceTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCSYSTEM": `ENTITY IfcSystem
 SUPERTYPE OF (ONEOF
	(IfcBuildingSystem
	,IfcBuiltSystem
	,IfcDistributionSystem
	,IfcStructuralAnalysisModel
	,IfcZone))
 SUBTYPE OF (IfcGroup);
 INVERSE
	ServicesBuildings : SET [0:1] OF IfcRelServicesBuildings FOR RelatingSystem;
	ServicesFacilities : SET [0:?] OF IfcRelReferencedInSpatialStructure FOR RelatedElements;
END_ENTITY;`,
	"IFCSYSTEMFURNITUREELEMENT": `ENTITY IfcSystemFurnitureElement
 SUBTYPE OF (IfcFurnishingElement);
	PredefinedType : OPTIONAL IfcSystemFurnitureElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSystemFurnitureElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSystemFurnitureElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCSYSTEMFURNITUREELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCSYSTEMFURNITUREELEMENTTYPE": `ENTITY IfcSystemFurnitureElementType
 SUBTYPE OF (IfcFurnishingElementType);
	PredefinedType : OPTIONAL IfcSystemFurnitureElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcSystemFurnitureElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcSystemFurnitureElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTABLE": `ENTITY IfcTable;
	Name : OPTIONAL IfcLabel;
	Rows : OPTIONAL LIST [1:?] OF IfcTableRow;
	Columns : OPTIONAL LIST [1:?] OF IfcTableColumn;
 DERIVE
	 NumberOfCellsInRow : IfcInteger := HIINDEX(Rows[1].RowCells);
	 NumberOfHeadings : IfcInteger := SIZEOF(QUERY( Temp <* Rows | Temp.IsHeading));
	 NumberOfDataRows : IfcInteger := SIZEOF(QUERY( Temp <* Rows | NOT(Temp.IsHeading)));
 WHERE
	WR1 : SIZEOF(QUERY( Temp <* Rows | HIINDEX(Temp.RowCells) <> HIINDEX(Rows[1].RowCells))) = 0;
	WR2 : { 0 <= NumberOfHeadings <= 1 };
END_ENTITY;`,
	"IFCTABLECOLUMN": `ENTITY IfcTableColumn;
	Identifier : OPTIONAL IfcIdentifier;
	Name : OPTIONAL IfcLabel;
	Description : OPTIONAL IfcText;
	Unit : OPTIONAL IfcUnit;
	ReferencePath : OPTIONAL IfcReference;
END_ENTITY;`,
	"IFCTABLEROW": `ENTITY IfcTableRow;
	RowCells : OPTIONAL LIST [1:?] OF IfcValue;
	IsHeading : OPTIONAL IfcBoolean;
END_ENTITY;`,
	"IFCTANK": `ENTITY IfcTank
 SUBTYPE OF (IfcFlowStorageDevice);
	PredefinedType : OPTIONAL IfcTankTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTankTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTankTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTANKTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTANKTYPE": `ENTITY IfcTankType
 SUBTYPE OF (IfcFlowStorageDeviceType);
	PredefinedType : IfcTankTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTankTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTankTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTASK": `ENTITY IfcTask
 SUBTYPE OF (IfcProcess);
	Status : OPTIONAL IfcLabel;
	WorkMethod : OPTIONAL IfcLabel;
	IsMilestone : IfcBoolean;
	Priority : OPTIONAL IfcInteger;
	TaskTime : OPTIONAL IfcTaskTime;
	PredefinedType : OPTIONAL IfcTaskTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcTaskTypeEnum.USERDEFINED) OR ((PredefinedType = IfcTaskTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
	HasName : EXISTS(SELF\IfcRoot.Name);
END_ENTITY;`,
	"IFCTASKTIME": `ENTITY IfcTaskTime
 SUPERTYPE OF (ONEOF
	(IfcTaskTimeRecurring))
 SUBTYPE OF (IfcSchedulingTime);
	DurationType : OPTIONAL IfcTaskDurationEnum;
	ScheduleDuration : OPTIONAL IfcDuration;
	ScheduleStart : OPTIONAL IfcDateTime;
	ScheduleFinish : OPTIONAL IfcDateTime;
	EarlyStart : OPTIONAL IfcDateTime;
	EarlyFinish : OPTIONAL IfcDateTime;
	LateStart : OPTIONAL IfcDateTime;
	LateFinish : OPTIONAL IfcDateTime;
	FreeFloat : OPTIONAL IfcDuration;
	TotalFloat : OPTIONAL IfcDuration;
	IsCritical : OPTIONAL IfcBoolean;
	StatusTime : OPTIONAL IfcDateTime;
	ActualDuration : OPTIONAL IfcDuration;
	ActualStart : OPTIONAL IfcDateTime;
	ActualFinish : OPTIONAL IfcDateTime;
	RemainingTime : OPTIONAL IfcDuration;
	Completion : OPTIONAL IfcPositiveRatioMeasure;
END_ENTITY;`,
	"IFCTASKTIMERECURRING": `ENTITY IfcTaskTimeRecurring
 SUBTYPE OF (IfcTaskTime);
	Recurrence : IfcRecurrencePattern;
END_ENTITY;`,
	"IFCTASKTYPE": `ENTITY IfcTaskType
 SUBTYPE OF (IfcTypeProcess);
	PredefinedType : IfcTaskTypeEnum;
	WorkMethod : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTaskTypeEnum.USERDEFINED) OR ((PredefinedType = IfcTaskTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcTypeProcess.ProcessType));
END_ENTITY;`,
	"IFCTELECOMADDRESS": `ENTITY IfcTelecomAddress
 SUBTYPE OF (IfcAddress);
	TelephoneNumbers : OPTIONAL LIST [1:?] OF IfcLabel;
	FacsimileNumbers : OPTIONAL LIST [1:?] OF IfcLabel;
	PagerNumber : OPTIONAL IfcLabel;
	ElectronicMailAddresses : OPTIONAL LIST [1:?] OF IfcLabel;
	WWWHomePageURL : OPTIONAL IfcURIReference;
	MessagingIDs : OPTIONAL LIST [1:?] OF IfcURIReference;
 WHERE
	MinimumDataProvided : EXISTS (TelephoneNumbers) OR
EXISTS (FacsimileNumbers) OR 
EXISTS (PagerNumber) OR
EXISTS (ElectronicMailAddresses) OR 
EXISTS (WWWHomePageURL) OR
EXISTS (MessagingIDs);
END_ENTITY;`,
	"IFCTENDON": `ENTITY IfcTendon
 SUBTYPE OF (IfcReinforcingElement);
	PredefinedType : OPTIONAL IfcTendonTypeEnum;
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	CrossSectionArea : OPTIONAL IfcAreaMeasure;
	TensionForce : OPTIONAL IfcForceMeasure;
	PreStress : OPTIONAL IfcPressureMeasure;
	FrictionCoefficient : OPTIONAL IfcNormalisedRatioMeasure;
	AnchorageSlip : OPTIONAL IfcPositiveLengthMeasure;
	MinCurvatureRadius : OPTIONAL IfcPositiveLengthMeasure;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTendonTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTENDONTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTENDONANCHOR": `ENTITY IfcTendonAnchor
 SUBTYPE OF (IfcReinforcingElement);
	PredefinedType : OPTIONAL IfcTendonAnchorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTendonAnchorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonAnchorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTENDONANCHORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTENDONANCHORTYPE": `ENTITY IfcTendonAnchorType
 SUBTYPE OF (IfcReinforcingElementType);
	PredefinedType : IfcTendonAnchorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTendonAnchorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonAnchorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTENDONCONDUIT": `ENTITY IfcTendonConduit
 SUBTYPE OF (IfcReinforcingElement);
	PredefinedType : OPTIONAL IfcTendonConduitTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTendonConduitTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonConduitTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTENDONCONDUITTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTENDONCONDUITTYPE": `ENTITY IfcTendonConduitType
 SUBTYPE OF (IfcReinforcingElementType);
	PredefinedType : IfcTendonConduitTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTendonConduitTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonConduitTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTENDONTYPE": `ENTITY IfcTendonType
 SUBTYPE OF (IfcReinforcingElementType);
	PredefinedType : IfcTendonTypeEnum;
	NominalDiameter : OPTIONAL IfcPositiveLengthMeasure;
	CrossSectionArea : OPTIONAL IfcAreaMeasure;
	SheathDiameter : OPTIONAL IfcPositiveLengthMeasure;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTendonTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTendonTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTESSELLATEDFACESET": `ENTITY IfcTessellatedFaceSet
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcPolygonalFaceSet
	,IfcTriangulatedFaceSet))
 SUBTYPE OF (IfcTessellatedItem);
	Coordinates : IfcCartesianPointList3D;
 DERIVE
	 Dim : IfcDimensionCount := 3;
 INVERSE
	HasColours : SET [0:1] OF IfcIndexedColourMap FOR MappedTo;
	HasTextures : SET [0:?] OF IfcIndexedTextureMap FOR MappedTo;
END_ENTITY;`,
	"IFCTESSELLATEDITEM": `ENTITY IfcTessellatedItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcIndexedPolygonalFace
	,IfcTessellatedFaceSet))
 SUBTYPE OF (IfcGeometricRepresentationItem);
END_ENTITY;`,
	"IFCTEXTLITERAL": `ENTITY IfcTextLiteral
 SUPERTYPE OF (ONEOF
	(IfcTextLiteralWithExtent))
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Literal : IfcPresentableText;
	Placement : IfcAxis2Placement;
	Path : IfcTextPath;
END_ENTITY;`,
	"IFCTEXTLITERALWITHEXTENT": `ENTITY IfcTextLiteralWithExtent
 SUBTYPE OF (IfcTextLiteral);
	Extent : IfcPlanarExtent;
	BoxAlignment : IfcBoxAlignment;
 WHERE
	WR31 : NOT('IFC4X3_DEV_73740fe4.IFCPLANARBOX' IN TYPEOF(Extent));
END_ENTITY;`,
	"IFCTEXTSTYLE": `ENTITY IfcTextStyle
 SUBTYPE OF (IfcPresentationStyle);
	TextCharacterAppearance : OPTIONAL IfcTextStyleForDefinedFont;
	TextStyle : OPTIONAL IfcTextStyleTextModel;
	TextFontStyle : IfcTextFontSelect;
	ModelOrDraughting : OPTIONAL IfcBoolean;
END_ENTITY;`,
	"IFCTEXTSTYLEFONTMODEL": `ENTITY IfcTextStyleFontModel
 SUBTYPE OF (IfcPreDefinedTextFont);
	FontFamily : LIST [1:?] OF IfcTextFontName;
	FontStyle : OPTIONAL IfcFontStyle;
	FontVariant : OPTIONAL IfcFontVariant;
	FontWeight : OPTIONAL IfcFontWeight;
	FontSize : IfcSizeSelect;
 WHERE
	MeasureOfFontSize : ('IFC4X3_DEV_73740fe4.IFCLENGTHMEASURE' IN TYPEOF(SELF.FontSize)) AND
(SELF.FontSize > 0.);
END_ENTITY;`,
	"IFCTEXTSTYLEFORDEFINEDFONT": `ENTITY IfcTextStyleForDefinedFont
 SUBTYPE OF (IfcPresentationItem);
	Colour : IfcColour;
	BackgroundColour : OPTIONAL IfcColour;
END_ENTITY;`,
	"IFCTEXTSTYLETEXTMODEL": `ENTITY IfcTextStyleTextModel
 SUBTYPE OF (IfcPresentationItem);
	TextIndent : OPTIONAL IfcSizeSelect;
	TextAlign : OPTIONAL IfcTextAlignment;
	TextDecoration : OPTIONAL IfcTextDecoration;
	LetterSpacing : OPTIONAL IfcSizeSelect;
	WordSpacing : OPTIONAL IfcSizeSelect;
	TextTransform : OPTIONAL IfcTextTransformation;
	LineHeight : OPTIONAL IfcSizeSelect;
END_ENTITY;`,
	"IFCTEXTURECOORDINATE": `ENTITY IfcTextureCoordinate
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcIndexedTextureMap
	,IfcTextureCoordinateGenerator
	,IfcTextureMap))
 SUBTYPE OF (IfcPresentationItem);
	Maps : LIST [1:?] OF IfcSurfaceTexture;
END_ENTITY;`,
	"IFCTEXTURECOORDINATEGENERATOR": `ENTITY IfcTextureCoordinateGenerator
 SUBTYPE OF (IfcTextureCoordinate);
	Mode : IfcLabel;
	Parameter : OPTIONAL LIST [1:?] OF IfcReal;
END_ENTITY;`,
	"IFCTEXTURECOORDINATEINDICES": `ENTITY IfcTextureCoordinateIndices
 SUPERTYPE OF (ONEOF
	(IfcTextureCoordinateIndicesWithVoids));
	TexCoordIndex : LIST [3:?] OF IfcPositiveInteger;
	TexCoordsOf : IfcIndexedPolygonalFace;
 INVERSE
	ToTexMap : IfcIndexedPolygonalTextureMap FOR TexCoordIndices;
END_ENTITY;`,
	"IFCTEXTURECOORDINATEINDICESWITHVOIDS": `ENTITY IfcTextureCoordinateIndicesWithVoids
 SUBTYPE OF (IfcTextureCoordinateIndices);
	InnerTexCoordIndices : LIST [1:?] OF LIST [3:?] OF UNIQUE IfcPositiveInteger;
END_ENTITY;`,
	"IFCTEXTUREMAP": `ENTITY IfcTextureMap
 SUBTYPE OF (IfcTextureCoordinate);
	Vertices : LIST [3:?] OF IfcTextureVertex;
	MappedTo : IfcFace;
END_ENTITY;`,
	"IFCTEXTUREVERTEX": `ENTITY IfcTextureVertex
 SUBTYPE OF (IfcPresentationItem);
	Coordinates : LIST [2:2] OF IfcParameterValue;
END_ENTITY;`,
	"IFCTEXTUREVERTEXLIST": `ENTITY IfcTextureVertexList
 SUBTYPE OF (IfcPresentationItem);
	TexCoordsList : LIST [1:?] OF LIST [2:2] OF IfcParameterValue;
END_ENTITY;`,
	"IFCTHIRDORDERPOLYNOMIALSPIRAL": `ENTITY IfcThirdOrderPolynomialSpiral
 SUBTYPE OF (IfcSpiral);
	CubicTerm : IfcLengthMeasure;
	QuadraticTerm : OPTIONAL IfcLengthMeasure;
	LinearTerm : OPTIONAL IfcLengthMeasure;
	ConstantTerm : OPTIONAL IfcLengthMeasure;
END_ENTITY;`,
	"IFCTIMEPERIOD": `ENTITY IfcTimePeriod;
	StartTime : IfcTime;
	EndTime : IfcTime;
END_ENTITY;`,
	"IFCTIMESERIES": `ENTITY IfcTimeSeries
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcIrregularTimeSeries
	,IfcRegularTimeSeries));
	Name : IfcLabel;
	Description : OPTIONAL IfcText;
	StartTime : IfcDateTime;
	EndTime : IfcDateTime;
	TimeSeriesDataType : IfcTimeSeriesDataTypeEnum;
	DataOrigin : IfcDataOriginEnum;
	UserDefinedDataOrigin : OPTIONAL IfcLabel;
	Unit : OPTIONAL IfcUnit;
 INVERSE
	HasExternalReference : SET [1:?] OF IfcExternalReferenceRelationship FOR RelatedResourceObjects;
END_ENTITY;`,
	"IFCTIMESERIESVALUE": `ENTITY IfcTimeSeriesValue;
	ListValues : LIST [1:?] OF IfcValue;
END_ENTITY;`,
	"IFCTOPOLOGICALREPRESENTATIONITEM": `ENTITY IfcTopologicalRepresentationItem
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConnectedFaceSet
	,IfcEdge
	,IfcFace
	,IfcFaceBound
	,IfcLoop
	,IfcPath
	,IfcVertex))
 SUBTYPE OF (IfcRepresentationItem);
END_ENTITY;`,
	"IFCTOPOLOGYREPRESENTATION": `ENTITY IfcTopologyRepresentation
 SUBTYPE OF (IfcShapeModel);
 WHERE
	WR21 : SIZEOF(QUERY(temp <* SELF\IfcRepresentation.Items | 
  NOT('IFC4X3_DEV_73740fe4.IFCTOPOLOGICALREPRESENTATIONITEM' IN TYPEOF(temp))
)) = 0

;
	WR22 : EXISTS(SELF\IfcRepresentation.RepresentationType);
	WR23 : IfcTopologyRepresentationTypes(SELF\IfcRepresentation.RepresentationType, SELF\IfcRepresentation.Items);
END_ENTITY;`,
	"IFCTOROIDALSURFACE": `ENTITY IfcToroidalSurface
 SUBTYPE OF (IfcElementarySurface);
	MajorRadius : IfcPositiveLengthMeasure;
	MinorRadius : IfcPositiveLengthMeasure;
 WHERE
	MajorLargerMinor : MinorRadius < MajorRadius;
END_ENTITY;`,
	"IFCTRACKELEMENT": `ENTITY IfcTrackElement
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcTrackElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTrackElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTrackElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTRACKELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTRACKELEMENTTYPE": `ENTITY IfcTrackElementType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcTrackElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTrackElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTrackElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTRANSFORMER": `ENTITY IfcTransformer
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcTransformerTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTransformerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTransformerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTRANSFORMERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTRANSFORMERTYPE": `ENTITY IfcTransformerType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcTransformerTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTransformerTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTransformerTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTRANSPORTATIONDEVICE": `ENTITY IfcTransportationDevice
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcTransportElement
	,IfcVehicle))
 SUBTYPE OF (IfcElement);
END_ENTITY;`,
	"IFCTRANSPORTATIONDEVICETYPE": `ENTITY IfcTransportationDeviceType
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcTransportElementType
	,IfcVehicleType))
 SUBTYPE OF (IfcElementType);
END_ENTITY;`,
	"IFCTRANSPORTELEMENT": `ENTITY IfcTransportElement
 SUBTYPE OF (IfcTransportationDevice);
	PredefinedType : OPTIONAL IfcTransportElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTransportElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTransportElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTRANSPORTELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTRANSPORTELEMENTTYPE": `ENTITY IfcTransportElementType
 SUBTYPE OF (IfcTransportationDeviceType);
	PredefinedType : IfcTransportElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTransportElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTransportElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTRAPEZIUMPROFILEDEF": `ENTITY IfcTrapeziumProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	BottomXDim : IfcPositiveLengthMeasure;
	TopXDim : IfcPositiveLengthMeasure;
	YDim : IfcPositiveLengthMeasure;
	TopXOffset : IfcLengthMeasure;
END_ENTITY;`,
	"IFCTRIANGULATEDFACESET": `ENTITY IfcTriangulatedFaceSet
 SUPERTYPE OF (ONEOF
	(IfcTriangulatedIrregularNetwork))
 SUBTYPE OF (IfcTessellatedFaceSet);
	Normals : OPTIONAL LIST [1:?] OF LIST [3:3] OF IfcParameterValue;
	Closed : OPTIONAL IfcBoolean;
	CoordIndex : LIST [1:?] OF LIST [3:3] OF IfcPositiveInteger;
	PnIndex : OPTIONAL LIST [1:?] OF IfcPositiveInteger;
 DERIVE
	 NumberOfTriangles : IfcInteger := SIZEOF(CoordIndex);
END_ENTITY;`,
	"IFCTRIANGULATEDIRREGULARNETWORK": `ENTITY IfcTriangulatedIrregularNetwork
 SUBTYPE OF (IfcTriangulatedFaceSet);
	Flags : LIST [1:?] OF IfcInteger;
 WHERE
	NotClosed : SELF\IfcTriangulatedFaceSet.Closed = FALSE;
END_ENTITY;`,
	"IFCTRIMMEDCURVE": `ENTITY IfcTrimmedCurve
 SUBTYPE OF (IfcBoundedCurve);
	BasisCurve : IfcCurve;
	Trim1 : SET [1:2] OF IfcTrimmingSelect;
	Trim2 : SET [1:2] OF IfcTrimmingSelect;
	SenseAgreement : IfcBoolean;
	MasterRepresentation : IfcTrimmingPreference;
 WHERE
	NoTrimOfBoundedCurves : NOT('IFC4X3_DEV_73740fe4.IFCBOUNDEDCURVE' IN TYPEOF(BasisCurve));
	Trim1ValuesConsistent : (HIINDEX(Trim1) = 1) OR (TYPEOF(Trim1[1]) <> TYPEOF(Trim1[2]));
	Trim2ValuesConsistent : (HIINDEX(Trim2) = 1) OR (TYPEOF(Trim2[1]) <> TYPEOF(Trim2[2]));
END_ENTITY;`,
	"IFCTSHAPEPROFILEDEF": `ENTITY IfcTShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	Depth : IfcPositiveLengthMeasure;
	FlangeWidth : IfcPositiveLengthMeasure;
	WebThickness : IfcPositiveLengthMeasure;
	FlangeThickness : IfcPositiveLengthMeasure;
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	FlangeEdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	WebEdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	WebSlope : OPTIONAL IfcPlaneAngleMeasure;
	FlangeSlope : OPTIONAL IfcPlaneAngleMeasure;
 WHERE
	ValidFlangeThickness : FlangeThickness < Depth;
	ValidWebThickness : WebThickness < FlangeWidth;
END_ENTITY;`,
	"IFCTUBEBUNDLE": `ENTITY IfcTubeBundle
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcTubeBundleTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcTubeBundleTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTubeBundleTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCTUBEBUNDLETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCTUBEBUNDLETYPE": `ENTITY IfcTubeBundleType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcTubeBundleTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcTubeBundleTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcTubeBundleTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCTYPEOBJECT": `ENTITY IfcTypeObject
 SUPERTYPE OF (ONEOF
	(IfcTypeProcess
	,IfcTypeProduct
	,IfcTypeResource))
 SUBTYPE OF (IfcObjectDefinition);
	ApplicableOccurrence : OPTIONAL IfcIdentifier;
	HasPropertySets : OPTIONAL SET [1:?] OF IfcPropertySetDefinition;
 INVERSE
	Types : SET [0:1] OF IfcRelDefinesByType FOR RelatingType;
 WHERE
	NameRequired : EXISTS(SELF\IfcRoot.Name);
	UniquePropertySetNames : (NOT(EXISTS(HasPropertySets))) OR IfcUniquePropertySetNames(HasPropertySets);
END_ENTITY;`,
	"IFCTYPEPROCESS": `ENTITY IfcTypeProcess
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcEventType
	,IfcProcedureType
	,IfcTaskType))
 SUBTYPE OF (IfcTypeObject);
	Identification : OPTIONAL IfcIdentifier;
	LongDescription : OPTIONAL IfcText;
	ProcessType : OPTIONAL IfcLabel;
 INVERSE
	OperatesOn : SET [0:?] OF IfcRelAssignsToProcess FOR RelatingProcess;
END_ENTITY;`,
	"IFCTYPEPRODUCT": `ENTITY IfcTypeProduct
 SUPERTYPE OF (ONEOF
	(IfcElementType
	,IfcSpatialElementType))
 SUBTYPE OF (IfcTypeObject);
	RepresentationMaps : OPTIONAL LIST [1:?] OF UNIQUE IfcRepresentationMap;
	Tag : OPTIONAL IfcLabel;
 INVERSE
	ReferencedBy : SET [0:?] OF IfcRelAssignsToProduct FOR RelatingProduct;
 WHERE
	ApplicableOccurrence : NOT(EXISTS(SELF\IfcTypeObject.Types[1])) OR
(SIZEOF(QUERY(temp <* SELF\IfcTypeObject.Types[1].RelatedObjects |
  NOT('IFC4X3_DEV_73740fe4.IFCPRODUCT' IN TYPEOF(temp)))
) = 0);
END_ENTITY;`,
	"IFCTYPERESOURCE": `ENTITY IfcTypeResource
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcConstructionResourceType))
 SUBTYPE OF (IfcTypeObject);
	Identification : OPTIONAL IfcIdentifier;
	LongDescription : OPTIONAL IfcText;
	ResourceType : OPTIONAL IfcLabel;
 INVERSE
	ResourceOf : SET [0:?] OF IfcRelAssignsToResource FOR RelatingResource;
END_ENTITY;`,
	"IFCUNITARYCONTROLELEMENT": `ENTITY IfcUnitaryControlElement
 SUBTYPE OF (IfcDistributionControlElement);
	PredefinedType : OPTIONAL IfcUnitaryControlElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcUnitaryControlElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcUnitaryControlElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCUNITARYCONTROLELEMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCUNITARYCONTROLELEMENTTYPE": `ENTITY IfcUnitaryControlElementType
 SUBTYPE OF (IfcDistributionControlElementType);
	PredefinedType : IfcUnitaryControlElementTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcUnitaryControlElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcUnitaryControlElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCUNITARYEQUIPMENT": `ENTITY IfcUnitaryEquipment
 SUBTYPE OF (IfcEnergyConversionDevice);
	PredefinedType : OPTIONAL IfcUnitaryEquipmentTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcUnitaryEquipmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcUnitaryEquipmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCUNITARYEQUIPMENTTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCUNITARYEQUIPMENTTYPE": `ENTITY IfcUnitaryEquipmentType
 SUBTYPE OF (IfcEnergyConversionDeviceType);
	PredefinedType : IfcUnitaryEquipmentTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcUnitaryEquipmentTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcUnitaryEquipmentTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCUNITASSIGNMENT": `ENTITY IfcUnitAssignment;
	Units : SET [1:?] OF IfcUnit;
 WHERE
	WR01 : IfcCorrectUnitAssignment(Units);
END_ENTITY;`,
	"IFCUSHAPEPROFILEDEF": `ENTITY IfcUShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	Depth : IfcPositiveLengthMeasure;
	FlangeWidth : IfcPositiveLengthMeasure;
	WebThickness : IfcPositiveLengthMeasure;
	FlangeThickness : IfcPositiveLengthMeasure;
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	EdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	FlangeSlope : OPTIONAL IfcPlaneAngleMeasure;
 WHERE
	ValidFlangeThickness : FlangeThickness < (Depth / 2.);
	ValidWebThickness : WebThickness < FlangeWidth;
END_ENTITY;`,
	"IFCVALVE": `ENTITY IfcValve
 SUBTYPE OF (IfcFlowController);
	PredefinedType : OPTIONAL IfcValveTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcValveTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcValveTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCVALVETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCVALVETYPE": `ENTITY IfcValveType
 SUBTYPE OF (IfcFlowControllerType);
	PredefinedType : IfcValveTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcValveTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcValveTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCVECTOR": `ENTITY IfcVector
 SUBTYPE OF (IfcGeometricRepresentationItem);
	Orientation : IfcDirection;
	Magnitude : IfcLengthMeasure;
 DERIVE
	 Dim : IfcDimensionCount := Orientation.Dim;
 WHERE
	MagGreaterOrEqualZero : Magnitude >= 0.0;
END_ENTITY;`,
	"IFCVEHICLE": `ENTITY IfcVehicle
 SUBTYPE OF (IfcTransportationDevice);
	PredefinedType : OPTIONAL IfcVehicleTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcVehicleTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVehicleTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCVEHICLETYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCVEHICLETYPE": `ENTITY IfcVehicleType
 SUBTYPE OF (IfcTransportationDeviceType);
	PredefinedType : IfcVehicleTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcVehicleTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVehicleTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCVERTEX": `ENTITY IfcVertex
 SUPERTYPE OF (ONEOF
	(IfcVertexPoint))
 SUBTYPE OF (IfcTopologicalRepresentationItem);
END_ENTITY;`,
	"IFCVERTEXLOOP": `ENTITY IfcVertexLoop
 SUBTYPE OF (IfcLoop);
	LoopVertex : IfcVertex;
END_ENTITY;`,
	"IFCVERTEXPOINT": `ENTITY IfcVertexPoint
 SUBTYPE OF (IfcVertex);
	VertexGeometry : IfcPoint;
END_ENTITY;`,
	"IFCVIBRATIONDAMPER": `ENTITY IfcVibrationDamper
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcVibrationDamperTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcVibrationDamperTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVibrationDamperTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCVIBRATIONDAMPERTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCVIBRATIONDAMPERTYPE": `ENTITY IfcVibrationDamperType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcVibrationDamperTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcVibrationDamperTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVibrationDamperTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCVIBRATIONISOLATOR": `ENTITY IfcVibrationIsolator
 SUBTYPE OF (IfcElementComponent);
	PredefinedType : OPTIONAL IfcVibrationIsolatorTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcVibrationIsolatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVibrationIsolatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCVIBRATIONISOLATORTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCVIBRATIONISOLATORTYPE": `ENTITY IfcVibrationIsolatorType
 SUBTYPE OF (IfcElementComponentType);
	PredefinedType : IfcVibrationIsolatorTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcVibrationIsolatorTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVibrationIsolatorTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCVIRTUALELEMENT": `ENTITY IfcVirtualElement
 SUBTYPE OF (IfcElement);
	PredefinedType : OPTIONAL IfcVirtualElementTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcVirtualElementTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVirtualElementTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCVIRTUALGRIDINTERSECTION": `ENTITY IfcVirtualGridIntersection;
	IntersectingAxes : LIST [2:2] OF UNIQUE IfcGridAxis;
	OffsetDistances : LIST [2:3] OF IfcLengthMeasure;
END_ENTITY;`,
	"IFCVOIDINGFEATURE": `ENTITY IfcVoidingFeature
 SUBTYPE OF (IfcFeatureElementSubtraction);
	PredefinedType : OPTIONAL IfcVoidingFeatureTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcVoidingFeatureTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcVoidingFeatureTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCWALL": `ENTITY IfcWall
 SUPERTYPE OF (ONEOF
	(IfcWallStandardCase))
 SUBTYPE OF (IfcBuiltElement);
	PredefinedType : OPTIONAL IfcWallTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcWallTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWallTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCWALLTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCWALLSTANDARDCASE": `ENTITY IfcWallStandardCase
 SUBTYPE OF (IfcWall);
 WHERE
	HasMaterialLayerSetUsage : SIZEOF (QUERY(temp <* USEDIN(SELF, 'IFC4X3_DEV_73740fe4.IFCRELASSOCIATES.RELATEDOBJECTS') |
              ('IFC4X3_DEV_73740fe4.IFCRELASSOCIATESMATERIAL' IN TYPEOF(temp)) AND
              ('IFC4X3_DEV_73740fe4.IFCMATERIALLAYERSETUSAGE' IN TYPEOF(temp.RelatingMaterial))
              )) = 1;
END_ENTITY;`,
	"IFCWALLTYPE": `ENTITY IfcWallType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcWallTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcWallTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWallTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCWASTETERMINAL": `ENTITY IfcWasteTerminal
 SUBTYPE OF (IfcFlowTerminal);
	PredefinedType : OPTIONAL IfcWasteTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcWasteTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWasteTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCWASTETERMINALTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCWASTETERMINALTYPE": `ENTITY IfcWasteTerminalType
 SUBTYPE OF (IfcFlowTerminalType);
	PredefinedType : IfcWasteTerminalTypeEnum;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcWasteTerminalTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWasteTerminalTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCWELLKNOWNTEXT": `ENTITY IfcWellKnownText;
	WellKnownText : IfcWellKnownTextLiteral;
	CoordinateReferenceSystem : IfcCoordinateReferenceSystem;
END_ENTITY;`,
	"IFCWINDOW": `ENTITY IfcWindow
 SUBTYPE OF (IfcBuiltElement);
	OverallHeight : OPTIONAL IfcPositiveLengthMeasure;
	OverallWidth : OPTIONAL IfcPositiveLengthMeasure;
	PredefinedType : OPTIONAL IfcWindowTypeEnum;
	PartitioningType : OPTIONAL IfcWindowTypePartitioningEnum;
	UserDefinedPartitioningType : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR
 (PredefinedType <> IfcWindowTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWindowTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcObject.ObjectType));
	CorrectTypeAssigned : (SIZEOF(IsTypedBy) = 0) OR
  ('IFC4X3_DEV_73740fe4.IFCWINDOWTYPE' IN TYPEOF(SELF\IfcObject.IsTypedBy[1].RelatingType));
END_ENTITY;`,
	"IFCWINDOWLININGPROPERTIES": `ENTITY IfcWindowLiningProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	LiningDepth : OPTIONAL IfcPositiveLengthMeasure;
	LiningThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	TransomThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	MullionThickness : OPTIONAL IfcNonNegativeLengthMeasure;
	FirstTransomOffset : OPTIONAL IfcNormalisedRatioMeasure;
	SecondTransomOffset : OPTIONAL IfcNormalisedRatioMeasure;
	FirstMullionOffset : OPTIONAL IfcNormalisedRatioMeasure;
	SecondMullionOffset : OPTIONAL IfcNormalisedRatioMeasure;
	ShapeAspectStyle : OPTIONAL IfcShapeAspect;
	LiningOffset : OPTIONAL IfcLengthMeasure;
	LiningToPanelOffsetX : OPTIONAL IfcLengthMeasure;
	LiningToPanelOffsetY : OPTIONAL IfcLengthMeasure;
 WHERE
	WR31 : NOT(EXISTS(LiningDepth) AND NOT(EXISTS(LiningThickness)));
	WR32 : NOT(NOT(EXISTS(FirstTransomOffset)) AND EXISTS(SecondTransomOffset));
	WR33 : NOT(NOT(EXISTS(FirstMullionOffset)) AND EXISTS(SecondMullionOffset));
	WR34 : (EXISTS(SELF\IfcPropertySetDefinition.DefinesType[1])) 
AND 
('IFC4X3_DEV_73740fe4.IFCWINDOWTYPE' IN TYPEOF(SELF\IfcPropertySetDefinition.DefinesType[1]));
END_ENTITY;`,
	"IFCWINDOWPANELPROPERTIES": `ENTITY IfcWindowPanelProperties
 SUBTYPE OF (IfcPreDefinedPropertySet);
	OperationType : IfcWindowPanelOperationEnum;
	PanelPosition : IfcWindowPanelPositionEnum;
	FrameDepth : OPTIONAL IfcPositiveLengthMeasure;
	FrameThickness : OPTIONAL IfcPositiveLengthMeasure;
	ShapeAspectStyle : OPTIONAL IfcShapeAspect;
 WHERE
	ApplicableToType : (EXISTS(SELF\IfcPropertySetDefinition.DefinesType[1])) 
AND 
('IFC4X3_DEV_73740fe4.IFCWINDOWTYPE' IN TYPEOF(SELF\IfcPropertySetDefinition.DefinesType[1]));
END_ENTITY;`,
	"IFCWINDOWTYPE": `ENTITY IfcWindowType
 SUBTYPE OF (IfcBuiltElementType);
	PredefinedType : IfcWindowTypeEnum;
	PartitioningType : IfcWindowTypePartitioningEnum;
	ParameterTakesPrecedence : OPTIONAL IfcBoolean;
	UserDefinedPartitioningType : OPTIONAL IfcLabel;
 WHERE
	CorrectPredefinedType : (PredefinedType <> IfcWindowTypeEnum.USERDEFINED) OR
 ((PredefinedType = IfcWindowTypeEnum.USERDEFINED) AND EXISTS (SELF\IfcElementType.ElementType));
END_ENTITY;`,
	"IFCWORKCALENDAR": `ENTITY IfcWorkCalendar
 SUBTYPE OF (IfcControl);
	WorkingTimes : OPTIONAL SET [1:?] OF IfcWorkTime;
	ExceptionTimes : OPTIONAL SET [1:?] OF IfcWorkTime;
	PredefinedType : OPTIONAL IfcWorkCalendarTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcWorkCalendarTypeEnum.USERDEFINED) OR
((PredefinedType = IfcWorkCalendarTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCWORKCONTROL": `ENTITY IfcWorkControl
 ABSTRACT SUPERTYPE OF (ONEOF
	(IfcWorkPlan
	,IfcWorkSchedule))
 SUBTYPE OF (IfcControl);
	CreationDate : IfcDateTime;
	Creators : OPTIONAL SET [1:?] OF IfcPerson;
	Purpose : OPTIONAL IfcLabel;
	Duration : OPTIONAL IfcDuration;
	TotalFloat : OPTIONAL IfcDuration;
	StartTime : IfcDateTime;
	FinishTime : OPTIONAL IfcDateTime;
END_ENTITY;`,
	"IFCWORKPLAN": `ENTITY IfcWorkPlan
 SUBTYPE OF (IfcWorkControl);
	PredefinedType : OPTIONAL IfcWorkPlanTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcWorkPlanTypeEnum.USERDEFINED) OR
((PredefinedType = IfcWorkPlanTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCWORKSCHEDULE": `ENTITY IfcWorkSchedule
 SUBTYPE OF (IfcWorkControl);
	PredefinedType : OPTIONAL IfcWorkScheduleTypeEnum;
 WHERE
	CorrectPredefinedType : NOT(EXISTS(PredefinedType)) OR (PredefinedType <> IfcWorkScheduleTypeEnum.USERDEFINED) OR
((PredefinedType = IfcWorkScheduleTypeEnum.USERDEFINED) AND EXISTS(SELF\IfcObject.ObjectType));
END_ENTITY;`,
	"IFCWORKTIME": `ENTITY IfcWorkTime
 SUBTYPE OF (IfcSchedulingTime);
	RecurrencePattern : OPTIONAL IfcRecurrencePattern;
	StartDate : OPTIONAL IfcDate;
	FinishDate : OPTIONAL IfcDate;
END_ENTITY;`,
	"IFCZONE": `ENTITY IfcZone
 SUBTYPE OF (IfcSystem);
	LongName : OPTIONAL IfcLabel;
 WHERE
	WR1 : (SIZEOF(SELF\IfcGroup.IsGroupedBy) = 0) OR
 (SIZEOF (QUERY (temp <* SELF\IfcGroup.IsGroupedBy[1].RelatedObjects |  
  NOT(('IFC4X3_DEV_73740fe4.IFCZONE' IN TYPEOF(temp)) OR 
  ('IFC4X3_DEV_73740fe4.IFCSPACE' IN TYPEOF(temp)) OR
  ('IFC4X3_DEV_73740fe4.IFCSPATIALZONE' IN TYPEOF(temp))
 ))) = 0);
END_ENTITY;`,
	"IFCZSHAPEPROFILEDEF": `ENTITY IfcZShapeProfileDef
 SUBTYPE OF (IfcParameterizedProfileDef);
	Depth : IfcPositiveLengthMeasure;
	FlangeWidth : IfcPositiveLengthMeasure;
	WebThickness : IfcPositiveLengthMeasure;
	FlangeThickness : IfcPositiveLengthMeasure;
	FilletRadius : OPTIONAL IfcNonNegativeLengthMeasure;
	EdgeRadius : OPTIONAL IfcNonNegativeLengthMeasure;
 WHERE
	ValidFlangeThickness : FlangeThickness < (Depth / 2.);
END_ENTITY;`,
}

var enumTypeExpressDocs = map[string]string{
	"IFCACTIONREQUESTTYPEENUM": `TYPE IfcActionRequestTypeEnum = ENUMERATION OF
	(EMAIL
	,FAX
	,PHONE
	,POST
	,VERBAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCACTIONSOURCETYPEENUM": `TYPE IfcActionSourceTypeEnum = ENUMERATION OF
	(BRAKES
	,BUOYANCY
	,COMPLETION_G1
	,CREEP
	,CURRENT
	,DEAD_LOAD_G
	,EARTHQUAKE_E
	,ERECTION
	,FIRE
	,ICE
	,IMPACT
	,IMPULSE
	,LACK_OF_FIT
	,LIVE_LOAD_Q
	,PRESTRESSING_P
	,PROPPING
	,RAIN
	,SETTLEMENT_U
	,SHRINKAGE
	,SNOW_S
	,SYSTEM_IMPERFECTION
	,TEMPERATURE_T
	,TRANSPORT
	,WAVE
	,WIND_W
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCACTIONTYPEENUM": `TYPE IfcActionTypeEnum = ENUMERATION OF
	(EXTRAORDINARY_A
	,PERMANENT_G
	,VARIABLE_Q
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCACTUATORTYPEENUM": `TYPE IfcActuatorTypeEnum = ENUMERATION OF
	(ELECTRICACTUATOR
	,HANDOPERATEDACTUATOR
	,HYDRAULICACTUATOR
	,PNEUMATICACTUATOR
	,THERMOSTATICACTUATOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCADDRESSTYPEENUM": `TYPE IfcAddressTypeEnum = ENUMERATION OF
	(DISTRIBUTIONPOINT
	,HOME
	,OFFICE
	,SITE
	,USERDEFINED);
END_TYPE;`,
	"IFCAIRTERMINALBOXTYPEENUM": `TYPE IfcAirTerminalBoxTypeEnum = ENUMERATION OF
	(CONSTANTFLOW
	,VARIABLEFLOWPRESSUREDEPENDANT
	,VARIABLEFLOWPRESSUREINDEPENDANT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCAIRTERMINALTYPEENUM": `TYPE IfcAirTerminalTypeEnum = ENUMERATION OF
	(DIFFUSER
	,GRILLE
	,LOUVRE
	,REGISTER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCAIRTOAIRHEATRECOVERYTYPEENUM": `TYPE IfcAirToAirHeatRecoveryTypeEnum = ENUMERATION OF
	(FIXEDPLATECOUNTERFLOWEXCHANGER
	,FIXEDPLATECROSSFLOWEXCHANGER
	,FIXEDPLATEPARALLELFLOWEXCHANGER
	,HEATPIPE
	,ROTARYWHEEL
	,RUNAROUNDCOILLOOP
	,THERMOSIPHONCOILTYPEHEATEXCHANGERS
	,THERMOSIPHONSEALEDTUBEHEATEXCHANGERS
	,TWINTOWERENTHALPYRECOVERYLOOPS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCALARMTYPEENUM": `TYPE IfcAlarmTypeEnum = ENUMERATION OF
	(BELL
	,BREAKGLASSBUTTON
	,LIGHT
	,MANUALPULLBOX
	,RAILWAYCROCODILE
	,RAILWAYDETONATOR
	,SIREN
	,WHISTLE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCALIGNMENTCANTSEGMENTTYPEENUM": `TYPE IfcAlignmentCantSegmentTypeEnum = ENUMERATION OF
	(BLOSSCURVE
	,CONSTANTCANT
	,COSINECURVE
	,HELMERTCURVE
	,LINEARTRANSITION
	,SINECURVE
	,VIENNESEBEND);
END_TYPE;`,
	"IFCALIGNMENTHORIZONTALSEGMENTTYPEENUM": `TYPE IfcAlignmentHorizontalSegmentTypeEnum = ENUMERATION OF
	(BLOSSCURVE
	,CIRCULARARC
	,CLOTHOID
	,COSINECURVE
	,CUBIC
	,HELMERTCURVE
	,LINE
	,SINECURVE
	,VIENNESEBEND);
END_TYPE;`,
	"IFCALIGNMENTTYPEENUM": `TYPE IfcAlignmentTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCALIGNMENTVERTICALSEGMENTTYPEENUM": `TYPE IfcAlignmentVerticalSegmentTypeEnum = ENUMERATION OF
	(CIRCULARARC
	,CLOTHOID
	,CONSTANTGRADIENT
	,PARABOLICARC);
END_TYPE;`,
	"IFCANALYSISMODELTYPEENUM": `TYPE IfcAnalysisModelTypeEnum = ENUMERATION OF
	(IN_PLANE_LOADING_2D
	,LOADING_3D
	,OUT_PLANE_LOADING_2D
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCANALYSISTHEORYTYPEENUM": `TYPE IfcAnalysisTheoryTypeEnum = ENUMERATION OF
	(FIRST_ORDER_THEORY
	,FULL_NONLINEAR_THEORY
	,SECOND_ORDER_THEORY
	,THIRD_ORDER_THEORY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCANNOTATIONTYPEENUM": `TYPE IfcAnnotationTypeEnum = ENUMERATION OF
	(CONTOURLINE
	,DIMENSION
	,ISOBAR
	,ISOLUX
	,ISOTHERM
	,LEADER
	,SURVEY
	,SYMBOL
	,TEXT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCARITHMETICOPERATORENUM": `TYPE IfcArithmeticOperatorEnum = ENUMERATION OF
	(ADD
	,DIVIDE
	,MODULO
	,MULTIPLY
	,SUBTRACT);
END_TYPE;`,
	"IFCASSEMBLYPLACEENUM": `TYPE IfcAssemblyPlaceEnum = ENUMERATION OF
	(FACTORY
	,SITE
	,NOTDEFINED);
END_TYPE;`,
	"IFCAUDIOVISUALAPPLIANCETYPEENUM": `TYPE IfcAudioVisualApplianceTypeEnum = ENUMERATION OF
	(AMPLIFIER
	,CAMERA
	,COMMUNICATIONTERMINAL
	,DISPLAY
	,MICROPHONE
	,PLAYER
	,PROJECTOR
	,RECEIVER
	,RECORDINGEQUIPMENT
	,SPEAKER
	,SWITCHER
	,TELEPHONE
	,TUNER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBEAMTYPEENUM": `TYPE IfcBeamTypeEnum = ENUMERATION OF
	(BEAM
	,CORNICE
	,DIAPHRAGM
	,EDGEBEAM
	,GIRDER_SEGMENT
	,HATSTONE
	,HOLLOWCORE
	,JOIST
	,LINTEL
	,PIERCAP
	,SPANDREL
	,T_BEAM
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBEARINGTYPEENUM": `TYPE IfcBearingTypeEnum = ENUMERATION OF
	(CYLINDRICAL
	,DISK
	,ELASTOMERIC
	,GUIDE
	,POT
	,ROCKER
	,ROLLER
	,SPHERICAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBENCHMARKENUM": `TYPE IfcBenchmarkEnum = ENUMERATION OF
	(EQUALTO
	,GREATERTHAN
	,GREATERTHANOREQUALTO
	,INCLUDEDIN
	,INCLUDES
	,LESSTHAN
	,LESSTHANOREQUALTO
	,NOTEQUALTO
	,NOTINCLUDEDIN
	,NOTINCLUDES);
END_TYPE;`,
	"IFCBOILERTYPEENUM": `TYPE IfcBoilerTypeEnum = ENUMERATION OF
	(STEAM
	,WATER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBOOLEANOPERATOR": `TYPE IfcBooleanOperator = ENUMERATION OF
	(DIFFERENCE
	,INTERSECTION
	,UNION);
END_TYPE;`,
	"IFCBRIDGEPARTTYPEENUM": `TYPE IfcBridgePartTypeEnum = ENUMERATION OF
	(ABUTMENT
	,DECK
	,DECK_SEGMENT
	,FOUNDATION
	,PIER
	,PIER_SEGMENT
	,PYLON
	,SUBSTRUCTURE
	,SUPERSTRUCTURE
	,SURFACESTRUCTURE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBRIDGETYPEENUM": `TYPE IfcBridgeTypeEnum = ENUMERATION OF
	(ARCHED
	,CABLE_STAYED
	,CANTILEVER
	,CULVERT
	,FRAMEWORK
	,GIRDER
	,SUSPENSION
	,TRUSS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBSPLINECURVEFORM": `TYPE IfcBSplineCurveForm = ENUMERATION OF
	(CIRCULAR_ARC
	,ELLIPTIC_ARC
	,HYPERBOLIC_ARC
	,PARABOLIC_ARC
	,POLYLINE_FORM
	,UNSPECIFIED);
END_TYPE;`,
	"IFCBSPLINESURFACEFORM": `TYPE IfcBSplineSurfaceForm = ENUMERATION OF
	(CONICAL_SURF
	,CYLINDRICAL_SURF
	,GENERALISED_CONE
	,PLANE_SURF
	,QUADRIC_SURF
	,RULED_SURF
	,SPHERICAL_SURF
	,SURF_OF_LINEAR_EXTRUSION
	,SURF_OF_REVOLUTION
	,TOROIDAL_SURF
	,UNSPECIFIED);
END_TYPE;`,
	"IFCBUILDINGELEMENTPARTTYPEENUM": `TYPE IfcBuildingElementPartTypeEnum = ENUMERATION OF
	(APRON
	,ARMOURUNIT
	,INSULATION
	,PRECASTPANEL
	,SAFETYCAGE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBUILDINGELEMENTPROXYTYPEENUM": `TYPE IfcBuildingElementProxyTypeEnum = ENUMERATION OF
	(COMPLEX
	,ELEMENT
	,PARTIAL
	,PROVISIONFORSPACE
	,PROVISIONFORVOID
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBUILDINGSYSTEMTYPEENUM": `TYPE IfcBuildingSystemTypeEnum = ENUMERATION OF
	(FENESTRATION
	,FOUNDATION
	,LOADBEARING
	,OUTERSHELL
	,SHADING
	,TRANSPORT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBUILTSYSTEMTYPEENUM": `TYPE IfcBuiltSystemTypeEnum = ENUMERATION OF
	(EROSIONPREVENTION
	,FENESTRATION
	,FOUNDATION
	,LOADBEARING
	,MOORING
	,OUTERSHELL
	,PRESTRESSING
	,RAILWAYLINE
	,RAILWAYTRACK
	,REINFORCING
	,SHADING
	,TRACKCIRCUIT
	,TRANSPORT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCBURNERTYPEENUM": `TYPE IfcBurnerTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCABLECARRIERFITTINGTYPEENUM": `TYPE IfcCableCarrierFittingTypeEnum = ENUMERATION OF
	(BEND
	,CONNECTOR
	,CROSS
	,JUNCTION
	,REDUCER
	,TEE
	,TRANSITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCABLECARRIERSEGMENTTYPEENUM": `TYPE IfcCableCarrierSegmentTypeEnum = ENUMERATION OF
	(CABLEBRACKET
	,CABLELADDERSEGMENT
	,CABLETRAYSEGMENT
	,CABLETRUNKINGSEGMENT
	,CATENARYWIRE
	,CONDUITSEGMENT
	,DROPPER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCABLEFITTINGTYPEENUM": `TYPE IfcCableFittingTypeEnum = ENUMERATION OF
	(CONNECTOR
	,ENTRY
	,EXIT
	,FANOUT
	,JUNCTION
	,TRANSITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCABLESEGMENTTYPEENUM": `TYPE IfcCableSegmentTypeEnum = ENUMERATION OF
	(BUSBARSEGMENT
	,CABLESEGMENT
	,CONDUCTORSEGMENT
	,CONTACTWIRESEGMENT
	,CORESEGMENT
	,FIBERSEGMENT
	,FIBERTUBE
	,OPTICALCABLESEGMENT
	,STITCHWIRE
	,WIREPAIRSEGMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCAISSONFOUNDATIONTYPEENUM": `TYPE IfcCaissonFoundationTypeEnum = ENUMERATION OF
	(CAISSON
	,WELL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCHANGEACTIONENUM": `TYPE IfcChangeActionEnum = ENUMERATION OF
	(ADDED
	,DELETED
	,MODIFIED
	,NOCHANGE
	,NOTDEFINED);
END_TYPE;`,
	"IFCCHILLERTYPEENUM": `TYPE IfcChillerTypeEnum = ENUMERATION OF
	(AIRCOOLED
	,HEATRECOVERY
	,WATERCOOLED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCHIMNEYTYPEENUM": `TYPE IfcChimneyTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOILTYPEENUM": `TYPE IfcCoilTypeEnum = ENUMERATION OF
	(DXCOOLINGCOIL
	,ELECTRICHEATINGCOIL
	,GASHEATINGCOIL
	,HYDRONICCOIL
	,STEAMHEATINGCOIL
	,WATERCOOLINGCOIL
	,WATERHEATINGCOIL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOLUMNTYPEENUM": `TYPE IfcColumnTypeEnum = ENUMERATION OF
	(COLUMN
	,PIERSTEM
	,PIERSTEM_SEGMENT
	,PILASTER
	,STANDCOLUMN
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOMMUNICATIONSAPPLIANCETYPEENUM": `TYPE IfcCommunicationsApplianceTypeEnum = ENUMERATION OF
	(ANTENNA
	,AUTOMATON
	,COMPUTER
	,FAX
	,GATEWAY
	,INTELLIGENTPERIPHERAL
	,IPNETWORKEQUIPMENT
	,LINESIDEELECTRONICUNIT
	,MODEM
	,NETWORKAPPLIANCE
	,NETWORKBRIDGE
	,NETWORKHUB
	,OPTICALLINETERMINAL
	,OPTICALNETWORKUNIT
	,PRINTER
	,RADIOBLOCKCENTER
	,REPEATER
	,ROUTER
	,SCANNER
	,TELECOMMAND
	,TELEPHONYEXCHANGE
	,TRANSITIONCOMPONENT
	,TRANSPONDER
	,TRANSPORTEQUIPMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOMPLEXPROPERTYTEMPLATETYPEENUM": `TYPE IfcComplexPropertyTemplateTypeEnum = ENUMERATION OF
	(P_COMPLEX
	,Q_COMPLEX);
END_TYPE;`,
	"IFCCOMPRESSORTYPEENUM": `TYPE IfcCompressorTypeEnum = ENUMERATION OF
	(BOOSTER
	,DYNAMIC
	,HERMETIC
	,OPENTYPE
	,RECIPROCATING
	,ROLLINGPISTON
	,ROTARY
	,ROTARYVANE
	,SCROLL
	,SEMIHERMETIC
	,SINGLESCREW
	,SINGLESTAGE
	,TROCHOIDAL
	,TWINSCREW
	,WELDEDSHELLHERMETIC
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONDENSERTYPEENUM": `TYPE IfcCondenserTypeEnum = ENUMERATION OF
	(AIRCOOLED
	,EVAPORATIVECOOLED
	,WATERCOOLED
	,WATERCOOLEDBRAZEDPLATE
	,WATERCOOLEDSHELLCOIL
	,WATERCOOLEDSHELLTUBE
	,WATERCOOLEDTUBEINTUBE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONNECTIONTYPEENUM": `TYPE IfcConnectionTypeEnum = ENUMERATION OF
	(ATEND
	,ATPATH
	,ATSTART
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONSTRAINTENUM": `TYPE IfcConstraintEnum = ENUMERATION OF
	(ADVISORY
	,HARD
	,SOFT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONSTRUCTIONEQUIPMENTRESOURCETYPEENUM": `TYPE IfcConstructionEquipmentResourceTypeEnum = ENUMERATION OF
	(DEMOLISHING
	,EARTHMOVING
	,ERECTING
	,HEATING
	,LIGHTING
	,PAVING
	,PUMPING
	,TRANSPORTING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONSTRUCTIONMATERIALRESOURCETYPEENUM": `TYPE IfcConstructionMaterialResourceTypeEnum = ENUMERATION OF
	(AGGREGATES
	,CONCRETE
	,DRYWALL
	,FUEL
	,GYPSUM
	,MASONRY
	,METAL
	,PLASTIC
	,WOOD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONSTRUCTIONPRODUCTRESOURCETYPEENUM": `TYPE IfcConstructionProductResourceTypeEnum = ENUMERATION OF
	(ASSEMBLY
	,FORMWORK
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONTROLLERTYPEENUM": `TYPE IfcControllerTypeEnum = ENUMERATION OF
	(FLOATING
	,MULTIPOSITION
	,PROGRAMMABLE
	,PROPORTIONAL
	,TWOPOSITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCONVEYORSEGMENTTYPEENUM": `TYPE IfcConveyorSegmentTypeEnum = ENUMERATION OF
	(BELTCONVEYOR
	,BUCKETCONVEYOR
	,CHUTECONVEYOR
	,SCREWCONVEYOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOOLEDBEAMTYPEENUM": `TYPE IfcCooledBeamTypeEnum = ENUMERATION OF
	(ACTIVE
	,PASSIVE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOOLINGTOWERTYPEENUM": `TYPE IfcCoolingTowerTypeEnum = ENUMERATION OF
	(MECHANICALFORCEDDRAFT
	,MECHANICALINDUCEDDRAFT
	,NATURALDRAFT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOSTITEMTYPEENUM": `TYPE IfcCostItemTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOSTSCHEDULETYPEENUM": `TYPE IfcCostScheduleTypeEnum = ENUMERATION OF
	(BUDGET
	,COSTPLAN
	,ESTIMATE
	,PRICEDBILLOFQUANTITIES
	,SCHEDULEOFRATES
	,TENDER
	,UNPRICEDBILLOFQUANTITIES
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOURSETYPEENUM": `TYPE IfcCourseTypeEnum = ENUMERATION OF
	(ARMOUR
	,BALLASTBED
	,CORE
	,FILTER
	,PAVEMENT
	,PROTECTION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCOVERINGTYPEENUM": `TYPE IfcCoveringTypeEnum = ENUMERATION OF
	(CEILING
	,CLADDING
	,COPING
	,FLOORING
	,INSULATION
	,MEMBRANE
	,MOLDING
	,ROOFING
	,SKIRTINGBOARD
	,SLEEVING
	,TOPPING
	,WRAPPING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCREWRESOURCETYPEENUM": `TYPE IfcCrewResourceTypeEnum = ENUMERATION OF
	(OFFICE
	,SITE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCURTAINWALLTYPEENUM": `TYPE IfcCurtainWallTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCCURVEINTERPOLATIONENUM": `TYPE IfcCurveInterpolationEnum = ENUMERATION OF
	(LINEAR
	,LOG_LINEAR
	,LOG_LOG
	,NOTDEFINED);
END_TYPE;`,
	"IFCDAMPERTYPEENUM": `TYPE IfcDamperTypeEnum = ENUMERATION OF
	(BACKDRAFTDAMPER
	,BALANCINGDAMPER
	,BLASTDAMPER
	,CONTROLDAMPER
	,FIREDAMPER
	,FIRESMOKEDAMPER
	,FUMEHOODEXHAUST
	,GRAVITYDAMPER
	,GRAVITYRELIEFDAMPER
	,RELIEFDAMPER
	,SMOKEDAMPER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDATAORIGINENUM": `TYPE IfcDataOriginEnum = ENUMERATION OF
	(MEASURED
	,PREDICTED
	,SIMULATED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDERIVEDUNITENUM": `TYPE IfcDerivedUnitEnum = ENUMERATION OF
	(ACCELERATIONUNIT
	,ANGULARVELOCITYUNIT
	,AREADENSITYUNIT
	,COMPOUNDPLANEANGLEUNIT
	,CURVATUREUNIT
	,DYNAMICVISCOSITYUNIT
	,HEATFLUXDENSITYUNIT
	,HEATINGVALUEUNIT
	,INTEGERCOUNTRATEUNIT
	,IONCONCENTRATIONUNIT
	,ISOTHERMALMOISTURECAPACITYUNIT
	,KINEMATICVISCOSITYUNIT
	,LINEARFORCEUNIT
	,LINEARMOMENTUNIT
	,LINEARSTIFFNESSUNIT
	,LINEARVELOCITYUNIT
	,LUMINOUSINTENSITYDISTRIBUTIONUNIT
	,MASSDENSITYUNIT
	,MASSFLOWRATEUNIT
	,MASSPERLENGTHUNIT
	,MODULUSOFELASTICITYUNIT
	,MODULUSOFLINEARSUBGRADEREACTIONUNIT
	,MODULUSOFROTATIONALSUBGRADEREACTIONUNIT
	,MODULUSOFSUBGRADEREACTIONUNIT
	,MOISTUREDIFFUSIVITYUNIT
	,MOLECULARWEIGHTUNIT
	,MOMENTOFINERTIAUNIT
	,PHUNIT
	,PLANARFORCEUNIT
	,ROTATIONALFREQUENCYUNIT
	,ROTATIONALMASSUNIT
	,ROTATIONALSTIFFNESSUNIT
	,SECTIONAREAINTEGRALUNIT
	,SECTIONMODULUSUNIT
	,SHEARMODULUSUNIT
	,SOUNDPOWERLEVELUNIT
	,SOUNDPOWERUNIT
	,SOUNDPRESSURELEVELUNIT
	,SOUNDPRESSUREUNIT
	,SPECIFICHEATCAPACITYUNIT
	,TEMPERATUREGRADIENTUNIT
	,TEMPERATURERATEOFCHANGEUNIT
	,THERMALADMITTANCEUNIT
	,THERMALCONDUCTANCEUNIT
	,THERMALEXPANSIONCOEFFICIENTUNIT
	,THERMALRESISTANCEUNIT
	,THERMALTRANSMITTANCEUNIT
	,TORQUEUNIT
	,VAPORPERMEABILITYUNIT
	,VOLUMETRICFLOWRATEUNIT
	,WARPINGCONSTANTUNIT
	,WARPINGMOMENTUNIT
	,USERDEFINED);
END_TYPE;`,
	"IFCDIRECTIONSENSEENUM": `TYPE IfcDirectionSenseEnum = ENUMERATION OF
	(NEGATIVE
	,POSITIVE);
END_TYPE;`,
	"IFCDISCRETEACCESSORYTYPEENUM": `TYPE IfcDiscreteAccessoryTypeEnum = ENUMERATION OF
	(ANCHORPLATE
	,BIRDPROTECTION
	,BRACKET
	,CABLEARRANGER
	,ELASTIC_CUSHION
	,EXPANSION_JOINT_DEVICE
	,FILLER
	,FLASHING
	,INSULATOR
	,LOCK
	,PANEL_STRENGTHENING
	,POINTMACHINEMOUNTINGDEVICE
	,POINT_MACHINE_LOCKING_DEVICE
	,RAILBRACE
	,RAILPAD
	,RAIL_LUBRICATION
	,RAIL_MECHANICAL_EQUIPMENT
	,SHOE
	,SLIDINGCHAIR
	,SOUNDABSORPTION
	,TENSIONINGEQUIPMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDISTRIBUTIONBOARDTYPEENUM": `TYPE IfcDistributionBoardTypeEnum = ENUMERATION OF
	(CONSUMERUNIT
	,DISPATCHINGBOARD
	,DISTRIBUTIONBOARD
	,DISTRIBUTIONFRAME
	,MOTORCONTROLCENTRE
	,SWITCHBOARD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDISTRIBUTIONCHAMBERELEMENTTYPEENUM": `TYPE IfcDistributionChamberElementTypeEnum = ENUMERATION OF
	(FORMEDDUCT
	,INSPECTIONCHAMBER
	,INSPECTIONPIT
	,MANHOLE
	,METERCHAMBER
	,SUMP
	,TRENCH
	,VALVECHAMBER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDISTRIBUTIONPORTTYPEENUM": `TYPE IfcDistributionPortTypeEnum = ENUMERATION OF
	(CABLE
	,CABLECARRIER
	,DUCT
	,PIPE
	,WIRELESS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDISTRIBUTIONSYSTEMENUM": `TYPE IfcDistributionSystemEnum = ENUMERATION OF
	(AIRCONDITIONING
	,AUDIOVISUAL
	,CATENARY_SYSTEM
	,CHEMICAL
	,CHILLEDWATER
	,COMMUNICATION
	,COMPRESSEDAIR
	,CONDENSERWATER
	,CONTROL
	,CONVEYING
	,DATA
	,DISPOSAL
	,DOMESTICCOLDWATER
	,DOMESTICHOTWATER
	,DRAINAGE
	,EARTHING
	,ELECTRICAL
	,ELECTROACOUSTIC
	,EXHAUST
	,FIREPROTECTION
	,FIXEDTRANSMISSIONNETWORK
	,FUEL
	,GAS
	,HAZARDOUS
	,HEATING
	,LIGHTING
	,LIGHTNINGPROTECTION
	,MOBILENETWORK
	,MONITORINGSYSTEM
	,MUNICIPALSOLIDWASTE
	,OIL
	,OPERATIONAL
	,OPERATIONALTELEPHONYSYSTEM
	,OVERHEAD_CONTACTLINE_SYSTEM
	,POWERGENERATION
	,RAINWATER
	,REFRIGERATION
	,RETURN_CIRCUIT
	,SECURITY
	,SEWAGE
	,SIGNAL
	,STORMWATER
	,TELEPHONE
	,TV
	,VACUUM
	,VENT
	,VENTILATION
	,WASTEWATER
	,WATERSUPPLY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOCUMENTCONFIDENTIALITYENUM": `TYPE IfcDocumentConfidentialityEnum = ENUMERATION OF
	(CONFIDENTIAL
	,PERSONAL
	,PUBLIC
	,RESTRICTED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOCUMENTSTATUSENUM": `TYPE IfcDocumentStatusEnum = ENUMERATION OF
	(DRAFT
	,FINAL
	,FINALDRAFT
	,REVISION
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOORPANELOPERATIONENUM": `TYPE IfcDoorPanelOperationEnum = ENUMERATION OF
	(DOUBLE_ACTING
	,FIXEDPANEL
	,FOLDING
	,REVOLVING
	,ROLLINGUP
	,SLIDING
	,SWINGING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOORPANELPOSITIONENUM": `TYPE IfcDoorPanelPositionEnum = ENUMERATION OF
	(LEFT
	,MIDDLE
	,RIGHT
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOORTYPEENUM": `TYPE IfcDoorTypeEnum = ENUMERATION OF
	(BOOM_BARRIER
	,DOOR
	,GATE
	,TRAPDOOR
	,TURNSTILE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDOORTYPEOPERATIONENUM": `TYPE IfcDoorTypeOperationEnum = ENUMERATION OF
	(DOUBLE_DOOR_DOUBLE_SWING
	,DOUBLE_DOOR_FOLDING
	,DOUBLE_DOOR_LIFTING_VERTICAL
	,DOUBLE_DOOR_SINGLE_SWING
	,DOUBLE_DOOR_SINGLE_SWING_OPPOSITE_LEFT
	,DOUBLE_DOOR_SINGLE_SWING_OPPOSITE_RIGHT
	,DOUBLE_DOOR_SLIDING
	,DOUBLE_SWING_LEFT
	,DOUBLE_SWING_RIGHT
	,FOLDING_TO_LEFT
	,FOLDING_TO_RIGHT
	,LIFTING_HORIZONTAL
	,LIFTING_VERTICAL_LEFT
	,LIFTING_VERTICAL_RIGHT
	,REVOLVING
	,REVOLVING_VERTICAL
	,ROLLINGUP
	,SINGLE_SWING_LEFT
	,SINGLE_SWING_RIGHT
	,SLIDING_TO_LEFT
	,SLIDING_TO_RIGHT
	,SWING_FIXED_LEFT
	,SWING_FIXED_RIGHT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDUCTFITTINGTYPEENUM": `TYPE IfcDuctFittingTypeEnum = ENUMERATION OF
	(BEND
	,CONNECTOR
	,ENTRY
	,EXIT
	,JUNCTION
	,OBSTRUCTION
	,TRANSITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDUCTSEGMENTTYPEENUM": `TYPE IfcDuctSegmentTypeEnum = ENUMERATION OF
	(FLEXIBLESEGMENT
	,RIGIDSEGMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCDUCTSILENCERTYPEENUM": `TYPE IfcDuctSilencerTypeEnum = ENUMERATION OF
	(FLATOVAL
	,RECTANGULAR
	,ROUND
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEARTHWORKSCUTTYPEENUM": `TYPE IfcEarthworksCutTypeEnum = ENUMERATION OF
	(BASE_EXCAVATION
	,CUT
	,DREDGING
	,EXCAVATION
	,OVEREXCAVATION
	,PAVEMENTMILLING
	,STEPEXCAVATION
	,TOPSOILREMOVAL
	,TRENCH
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEARTHWORKSFILLTYPEENUM": `TYPE IfcEarthworksFillTypeEnum = ENUMERATION OF
	(BACKFILL
	,COUNTERWEIGHT
	,EMBANKMENT
	,SLOPEFILL
	,SUBGRADE
	,SUBGRADEBED
	,TRANSITIONSECTION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICAPPLIANCETYPEENUM": `TYPE IfcElectricApplianceTypeEnum = ENUMERATION OF
	(DISHWASHER
	,ELECTRICCOOKER
	,FREESTANDINGELECTRICHEATER
	,FREESTANDINGFAN
	,FREESTANDINGWATERCOOLER
	,FREESTANDINGWATERHEATER
	,FREEZER
	,FRIDGE_FREEZER
	,HANDDRYER
	,KITCHENMACHINE
	,MICROWAVE
	,PHOTOCOPIER
	,REFRIGERATOR
	,TUMBLEDRYER
	,VENDINGMACHINE
	,WASHINGMACHINE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICDISTRIBUTIONBOARDTYPEENUM": `TYPE IfcElectricDistributionBoardTypeEnum = ENUMERATION OF
	(CONSUMERUNIT
	,DISTRIBUTIONBOARD
	,MOTORCONTROLCENTRE
	,SWITCHBOARD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICFLOWSTORAGEDEVICETYPEENUM": `TYPE IfcElectricFlowStorageDeviceTypeEnum = ENUMERATION OF
	(BATTERY
	,CAPACITOR
	,CAPACITORBANK
	,COMPENSATOR
	,HARMONICFILTER
	,INDUCTOR
	,INDUCTORBANK
	,RECHARGER
	,UPS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICFLOWTREATMENTDEVICETYPEENUM": `TYPE IfcElectricFlowTreatmentDeviceTypeEnum = ENUMERATION OF
	(ELECTRONICFILTER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICGENERATORTYPEENUM": `TYPE IfcElectricGeneratorTypeEnum = ENUMERATION OF
	(CHP
	,ENGINEGENERATOR
	,STANDALONE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICMOTORTYPEENUM": `TYPE IfcElectricMotorTypeEnum = ENUMERATION OF
	(DC
	,INDUCTION
	,POLYPHASE
	,RELUCTANCESYNCHRONOUS
	,SYNCHRONOUS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELECTRICTIMECONTROLTYPEENUM": `TYPE IfcElectricTimeControlTypeEnum = ENUMERATION OF
	(RELAY
	,TIMECLOCK
	,TIMEDELAY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELEMENTASSEMBLYTYPEENUM": `TYPE IfcElementAssemblyTypeEnum = ENUMERATION OF
	(ABUTMENT
	,ACCESSORY_ASSEMBLY
	,ARCH
	,BEAM_GRID
	,BRACED_FRAME
	,CROSS_BRACING
	,DECK
	,DILATATIONPANEL
	,ENTRANCEWORKS
	,GIRDER
	,GRID
	,MAST
	,PIER
	,PYLON
	,RAIL_MECHANICAL_EQUIPMENT_ASSEMBLY
	,REINFORCEMENT_UNIT
	,RIGID_FRAME
	,SHELTER
	,SIGNALASSEMBLY
	,SLAB_FIELD
	,SUMPBUSTER
	,SUPPORTINGASSEMBLY
	,SUSPENSIONASSEMBLY
	,TRACKPANEL
	,TRACTION_SWITCHING_ASSEMBLY
	,TRAFFIC_CALMING_DEVICE
	,TRUSS
	,TURNOUTPANEL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCELEMENTCOMPOSITIONENUM": `TYPE IfcElementCompositionEnum = ENUMERATION OF
	(COMPLEX
	,ELEMENT
	,PARTIAL);
END_TYPE;`,
	"IFCENGINETYPEENUM": `TYPE IfcEngineTypeEnum = ENUMERATION OF
	(EXTERNALCOMBUSTION
	,INTERNALCOMBUSTION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEVAPORATIVECOOLERTYPEENUM": `TYPE IfcEvaporativeCoolerTypeEnum = ENUMERATION OF
	(DIRECTEVAPORATIVEAIRWASHER
	,DIRECTEVAPORATIVEPACKAGEDROTARYAIRCOOLER
	,DIRECTEVAPORATIVERANDOMMEDIAAIRCOOLER
	,DIRECTEVAPORATIVERIGIDMEDIAAIRCOOLER
	,DIRECTEVAPORATIVESLINGERSPACKAGEDAIRCOOLER
	,INDIRECTDIRECTCOMBINATION
	,INDIRECTEVAPORATIVECOOLINGTOWERORCOILCOOLER
	,INDIRECTEVAPORATIVEPACKAGEAIRCOOLER
	,INDIRECTEVAPORATIVEWETCOIL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEVAPORATORTYPEENUM": `TYPE IfcEvaporatorTypeEnum = ENUMERATION OF
	(DIRECTEXPANSION
	,DIRECTEXPANSIONBRAZEDPLATE
	,DIRECTEXPANSIONSHELLANDTUBE
	,DIRECTEXPANSIONTUBEINTUBE
	,FLOODEDSHELLANDTUBE
	,SHELLANDCOIL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEVENTTRIGGERTYPEENUM": `TYPE IfcEventTriggerTypeEnum = ENUMERATION OF
	(EVENTCOMPLEX
	,EVENTMESSAGE
	,EVENTRULE
	,EVENTTIME
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEVENTTYPEENUM": `TYPE IfcEventTypeEnum = ENUMERATION OF
	(ENDEVENT
	,INTERMEDIATEEVENT
	,STARTEVENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCEXTERNALSPATIALELEMENTTYPEENUM": `TYPE IfcExternalSpatialElementTypeEnum = ENUMERATION OF
	(EXTERNAL
	,EXTERNAL_EARTH
	,EXTERNAL_FIRE
	,EXTERNAL_WATER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFACILITYPARTCOMMONTYPEENUM": `TYPE IfcFacilityPartCommonTypeEnum = ENUMERATION OF
	(ABOVEGROUND
	,BELOWGROUND
	,JUNCTION
	,LEVELCROSSING
	,SEGMENT
	,SUBSTRUCTURE
	,SUPERSTRUCTURE
	,TERMINAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFACILITYUSAGEENUM": `TYPE IfcFacilityUsageEnum = ENUMERATION OF
	(LATERAL
	,LONGITUDINAL
	,REGION
	,VERTICAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFANTYPEENUM": `TYPE IfcFanTypeEnum = ENUMERATION OF
	(CENTRIFUGALAIRFOIL
	,CENTRIFUGALBACKWARDINCLINEDCURVED
	,CENTRIFUGALFORWARDCURVED
	,CENTRIFUGALRADIAL
	,PROPELLORAXIAL
	,TUBEAXIAL
	,VANEAXIAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFASTENERTYPEENUM": `TYPE IfcFastenerTypeEnum = ENUMERATION OF
	(GLUE
	,MORTAR
	,WELD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFILTERTYPEENUM": `TYPE IfcFilterTypeEnum = ENUMERATION OF
	(AIRPARTICLEFILTER
	,COMPRESSEDAIRFILTER
	,ODORFILTER
	,OILFILTER
	,STRAINER
	,WATERFILTER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFIRESUPPRESSIONTERMINALTYPEENUM": `TYPE IfcFireSuppressionTerminalTypeEnum = ENUMERATION OF
	(BREECHINGINLET
	,FIREHYDRANT
	,FIREMONITOR
	,HOSEREEL
	,SPRINKLER
	,SPRINKLERDEFLECTOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFLOWDIRECTIONENUM": `TYPE IfcFlowDirectionEnum = ENUMERATION OF
	(SINK
	,SOURCE
	,SOURCEANDSINK
	,NOTDEFINED);
END_TYPE;`,
	"IFCFLOWINSTRUMENTTYPEENUM": `TYPE IfcFlowInstrumentTypeEnum = ENUMERATION OF
	(AMMETER
	,COMBINED
	,FREQUENCYMETER
	,PHASEANGLEMETER
	,POWERFACTORMETER
	,PRESSUREGAUGE
	,THERMOMETER
	,VOLTMETER
	,VOLTMETER_PEAK
	,VOLTMETER_RMS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFLOWMETERTYPEENUM": `TYPE IfcFlowMeterTypeEnum = ENUMERATION OF
	(ENERGYMETER
	,GASMETER
	,OILMETER
	,WATERMETER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFOOTINGTYPEENUM": `TYPE IfcFootingTypeEnum = ENUMERATION OF
	(CAISSON_FOUNDATION
	,FOOTING_BEAM
	,PAD_FOOTING
	,PILE_CAP
	,STRIP_FOOTING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCFURNITURETYPEENUM": `TYPE IfcFurnitureTypeEnum = ENUMERATION OF
	(BED
	,CHAIR
	,DESK
	,FILECABINET
	,SHELF
	,SOFA
	,TABLE
	,TECHNICALCABINET
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCGEOGRAPHICELEMENTTYPEENUM": `TYPE IfcGeographicElementTypeEnum = ENUMERATION OF
	(SOIL_BORING_POINT
	,TERRAIN
	,VEGETATION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCGEOMETRICPROJECTIONENUM": `TYPE IfcGeometricProjectionEnum = ENUMERATION OF
	(ELEVATION_VIEW
	,GRAPH_VIEW
	,MODEL_VIEW
	,PLAN_VIEW
	,REFLECTED_PLAN_VIEW
	,SECTION_VIEW
	,SKETCH_VIEW
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCGEOTECHNICALSTRATUMTYPEENUM": `TYPE IfcGeotechnicalStratumTypeEnum = ENUMERATION OF
	(SOLID
	,VOID
	,WATER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCGLOBALORLOCALENUM": `TYPE IfcGlobalOrLocalEnum = ENUMERATION OF
	(GLOBAL_COORDS
	,LOCAL_COORDS);
END_TYPE;`,
	"IFCGRIDTYPEENUM": `TYPE IfcGridTypeEnum = ENUMERATION OF
	(IRREGULAR
	,RADIAL
	,RECTANGULAR
	,TRIANGULAR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCHEATEXCHANGERTYPEENUM": `TYPE IfcHeatExchangerTypeEnum = ENUMERATION OF
	(PLATE
	,SHELLANDTUBE
	,TURNOUTHEATING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCHUMIDIFIERTYPEENUM": `TYPE IfcHumidifierTypeEnum = ENUMERATION OF
	(ADIABATICAIRWASHER
	,ADIABATICATOMIZING
	,ADIABATICCOMPRESSEDAIRNOZZLE
	,ADIABATICPAN
	,ADIABATICRIGIDMEDIA
	,ADIABATICULTRASONIC
	,ADIABATICWETTEDELEMENT
	,ASSISTEDBUTANE
	,ASSISTEDELECTRIC
	,ASSISTEDNATURALGAS
	,ASSISTEDPROPANE
	,ASSISTEDSTEAM
	,STEAMINJECTION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCIMPACTPROTECTIONDEVICETYPEENUM": `TYPE IfcImpactProtectionDeviceTypeEnum = ENUMERATION OF
	(BUMPER
	,CRASHCUSHION
	,DAMPINGSYSTEM
	,FENDER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCINTERCEPTORTYPEENUM": `TYPE IfcInterceptorTypeEnum = ENUMERATION OF
	(CYCLONIC
	,GREASE
	,OIL
	,PETROL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCINTERNALOREXTERNALENUM": `TYPE IfcInternalOrExternalEnum = ENUMERATION OF
	(EXTERNAL
	,EXTERNAL_EARTH
	,EXTERNAL_FIRE
	,EXTERNAL_WATER
	,INTERNAL
	,NOTDEFINED);
END_TYPE;`,
	"IFCINVENTORYTYPEENUM": `TYPE IfcInventoryTypeEnum = ENUMERATION OF
	(ASSETINVENTORY
	,FURNITUREINVENTORY
	,SPACEINVENTORY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCJUNCTIONBOXTYPEENUM": `TYPE IfcJunctionBoxTypeEnum = ENUMERATION OF
	(DATA
	,POWER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCKERBTYPEENUM": `TYPE IfcKerbTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCKNOTTYPE": `TYPE IfcKnotType = ENUMERATION OF
	(PIECEWISE_BEZIER_KNOTS
	,QUASI_UNIFORM_KNOTS
	,UNIFORM_KNOTS
	,UNSPECIFIED);
END_TYPE;`,
	"IFCLABORRESOURCETYPEENUM": `TYPE IfcLaborResourceTypeEnum = ENUMERATION OF
	(ADMINISTRATION
	,CARPENTRY
	,CLEANING
	,CONCRETE
	,DRYWALL
	,ELECTRIC
	,FINISHING
	,FLOORING
	,GENERAL
	,HVAC
	,LANDSCAPING
	,MASONRY
	,PAINTING
	,PAVING
	,PLUMBING
	,ROOFING
	,SITEGRADING
	,STEELWORK
	,SURVEYING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCLAMPTYPEENUM": `TYPE IfcLampTypeEnum = ENUMERATION OF
	(COMPACTFLUORESCENT
	,FLUORESCENT
	,HALOGEN
	,HIGHPRESSUREMERCURY
	,HIGHPRESSURESODIUM
	,LED
	,METALHALIDE
	,OLED
	,TUNGSTENFILAMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCLAYERSETDIRECTIONENUM": `TYPE IfcLayerSetDirectionEnum = ENUMERATION OF
	(AXIS1
	,AXIS2
	,AXIS3);
END_TYPE;`,
	"IFCLIGHTDISTRIBUTIONCURVEENUM": `TYPE IfcLightDistributionCurveEnum = ENUMERATION OF
	(TYPE_A
	,TYPE_B
	,TYPE_C
	,NOTDEFINED);
END_TYPE;`,
	"IFCLIGHTEMISSIONSOURCEENUM": `TYPE IfcLightEmissionSourceEnum = ENUMERATION OF
	(COMPACTFLUORESCENT
	,FLUORESCENT
	,HIGHPRESSUREMERCURY
	,HIGHPRESSURESODIUM
	,LIGHTEMITTINGDIODE
	,LOWPRESSURESODIUM
	,LOWVOLTAGEHALOGEN
	,MAINVOLTAGEHALOGEN
	,METALHALIDE
	,TUNGSTENFILAMENT
	,NOTDEFINED);
END_TYPE;`,
	"IFCLIGHTFIXTURETYPEENUM": `TYPE IfcLightFixtureTypeEnum = ENUMERATION OF
	(DIRECTIONSOURCE
	,POINTSOURCE
	,SECURITYLIGHTING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCLIQUIDTERMINALTYPEENUM": `TYPE IfcLiquidTerminalTypeEnum = ENUMERATION OF
	(HOSEREEL
	,LOADINGARM
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCLOADGROUPTYPEENUM": `TYPE IfcLoadGroupTypeEnum = ENUMERATION OF
	(LOAD_CASE
	,LOAD_COMBINATION
	,LOAD_GROUP
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCLOGICALOPERATORENUM": `TYPE IfcLogicalOperatorEnum = ENUMERATION OF
	(LOGICALAND
	,LOGICALNOTAND
	,LOGICALNOTOR
	,LOGICALOR
	,LOGICALXOR);
END_TYPE;`,
	"IFCMARINEFACILITYTYPEENUM": `TYPE IfcMarineFacilityTypeEnum = ENUMERATION OF
	(BARRIERBEACH
	,BREAKWATER
	,CANAL
	,DRYDOCK
	,FLOATINGDOCK
	,HYDROLIFT
	,JETTY
	,LAUNCHRECOVERY
	,MARINEDEFENCE
	,NAVIGATIONALCHANNEL
	,PORT
	,QUAY
	,REVETMENT
	,SHIPLIFT
	,SHIPLOCK
	,SHIPYARD
	,SLIPWAY
	,WATERWAY
	,WATERWAYSHIPLIFT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMARINEPARTTYPEENUM": `TYPE IfcMarinePartTypeEnum = ENUMERATION OF
	(ABOVEWATERLINE
	,ANCHORAGE
	,APPROACHCHANNEL
	,BELOWWATERLINE
	,BERTHINGSTRUCTURE
	,CHAMBER
	,CILL_LEVEL
	,COPELEVEL
	,CORE
	,CREST
	,GATEHEAD
	,GUDINGSTRUCTURE
	,HIGHWATERLINE
	,LANDFIELD
	,LEEWARDSIDE
	,LOWWATERLINE
	,MANUFACTURING
	,NAVIGATIONALAREA
	,PROTECTION
	,SHIPTRANSFER
	,STORAGEAREA
	,VEHICLESERVICING
	,WATERFIELD
	,WEATHERSIDE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMECHANICALFASTENERTYPEENUM": `TYPE IfcMechanicalFastenerTypeEnum = ENUMERATION OF
	(ANCHORBOLT
	,BOLT
	,CHAIN
	,COUPLER
	,DOWEL
	,NAIL
	,NAILPLATE
	,RAILFASTENING
	,RAILJOINT
	,RIVET
	,ROPE
	,SCREW
	,SHEARCONNECTOR
	,STAPLE
	,STUDSHEARCONNECTOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMEDICALDEVICETYPEENUM": `TYPE IfcMedicalDeviceTypeEnum = ENUMERATION OF
	(AIRSTATION
	,FEEDAIRUNIT
	,OXYGENGENERATOR
	,OXYGENPLANT
	,VACUUMSTATION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMEMBERTYPEENUM": `TYPE IfcMemberTypeEnum = ENUMERATION OF
	(ARCH_SEGMENT
	,BRACE
	,CHORD
	,COLLAR
	,MEMBER
	,MULLION
	,PLATE
	,POST
	,PURLIN
	,RAFTER
	,STAY_CABLE
	,STIFFENING_RIB
	,STRINGER
	,STRUCTURALCABLE
	,STRUT
	,STUD
	,SUSPENDER
	,SUSPENSION_CABLE
	,TIEBAR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMOBILETELECOMMUNICATIONSAPPLIANCETYPEENUM": `TYPE IfcMobileTelecommunicationsApplianceTypeEnum = ENUMERATION OF
	(ACCESSPOINT
	,BASEBANDUNIT
	,BASETRANSCEIVERSTATION
	,E_UTRAN_NODE_B
	,GATEWAY_GPRS_SUPPORT_NODE
	,MASTERUNIT
	,MOBILESWITCHINGCENTER
	,MSCSERVER
	,PACKETCONTROLUNIT
	,REMOTERADIOUNIT
	,REMOTEUNIT
	,SERVICE_GPRS_SUPPORT_NODE
	,SUBSCRIBERSERVER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMOORINGDEVICETYPEENUM": `TYPE IfcMooringDeviceTypeEnum = ENUMERATION OF
	(BOLLARD
	,LINETENSIONER
	,MAGNETICDEVICE
	,MOORINGHOOKS
	,VACUUMDEVICE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCMOTORCONNECTIONTYPEENUM": `TYPE IfcMotorConnectionTypeEnum = ENUMERATION OF
	(BELTDRIVE
	,COUPLING
	,DIRECTDRIVE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCNAVIGATIONELEMENTTYPEENUM": `TYPE IfcNavigationElementTypeEnum = ENUMERATION OF
	(BEACON
	,BUOY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCOBJECTIVEENUM": `TYPE IfcObjectiveEnum = ENUMERATION OF
	(CODECOMPLIANCE
	,CODEWAIVER
	,DESIGNINTENT
	,EXTERNAL
	,HEALTHANDSAFETY
	,MERGECONFLICT
	,MODELVIEW
	,PARAMETER
	,REQUIREMENT
	,SPECIFICATION
	,TRIGGERCONDITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCOCCUPANTTYPEENUM": `TYPE IfcOccupantTypeEnum = ENUMERATION OF
	(ASSIGNEE
	,ASSIGNOR
	,LESSEE
	,LESSOR
	,LETTINGAGENT
	,OWNER
	,TENANT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCOPENINGELEMENTTYPEENUM": `TYPE IfcOpeningElementTypeEnum = ENUMERATION OF
	(OPENING
	,RECESS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCOUTLETTYPEENUM": `TYPE IfcOutletTypeEnum = ENUMERATION OF
	(AUDIOVISUALOUTLET
	,COMMUNICATIONSOUTLET
	,DATAOUTLET
	,POWEROUTLET
	,TELEPHONEOUTLET
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPAVEMENTTYPEENUM": `TYPE IfcPavementTypeEnum = ENUMERATION OF
	(FLEXIBLE
	,RIGID
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPERFORMANCEHISTORYTYPEENUM": `TYPE IfcPerformanceHistoryTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPERMEABLECOVERINGOPERATIONENUM": `TYPE IfcPermeableCoveringOperationEnum = ENUMERATION OF
	(GRILL
	,LOUVER
	,SCREEN
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPERMITTYPEENUM": `TYPE IfcPermitTypeEnum = ENUMERATION OF
	(ACCESS
	,BUILDING
	,WORK
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPHYSICALORVIRTUALENUM": `TYPE IfcPhysicalOrVirtualEnum = ENUMERATION OF
	(PHYSICAL
	,VIRTUAL
	,NOTDEFINED);
END_TYPE;`,
	"IFCPILECONSTRUCTIONENUM": `TYPE IfcPileConstructionEnum = ENUMERATION OF
	(CAST_IN_PLACE
	,COMPOSITE
	,PRECAST_CONCRETE
	,PREFAB_STEEL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPILETYPEENUM": `TYPE IfcPileTypeEnum = ENUMERATION OF
	(BORED
	,COHESION
	,DRIVEN
	,FRICTION
	,JETGROUTING
	,SUPPORT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPIPEFITTINGTYPEENUM": `TYPE IfcPipeFittingTypeEnum = ENUMERATION OF
	(BEND
	,CONNECTOR
	,ENTRY
	,EXIT
	,JUNCTION
	,OBSTRUCTION
	,TRANSITION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPIPESEGMENTTYPEENUM": `TYPE IfcPipeSegmentTypeEnum = ENUMERATION OF
	(CULVERT
	,FLEXIBLESEGMENT
	,GUTTER
	,RIGIDSEGMENT
	,SPOOL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPLATETYPEENUM": `TYPE IfcPlateTypeEnum = ENUMERATION OF
	(BASE_PLATE
	,COVER_PLATE
	,CURTAIN_PANEL
	,FLANGE_PLATE
	,GUSSET_PLATE
	,SHEET
	,SPLICE_PLATE
	,STIFFENER_PLATE
	,WEB_PLATE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPREFERREDSURFACECURVEREPRESENTATION": `TYPE IfcPreferredSurfaceCurveRepresentation = ENUMERATION OF
	(CURVE3D
	,PCURVE_S1
	,PCURVE_S2);
END_TYPE;`,
	"IFCPROCEDURETYPEENUM": `TYPE IfcProcedureTypeEnum = ENUMERATION OF
	(ADVICE_CAUTION
	,ADVICE_NOTE
	,ADVICE_WARNING
	,CALIBRATION
	,DIAGNOSTIC
	,SHUTDOWN
	,STARTUP
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPROFILETYPEENUM": `TYPE IfcProfileTypeEnum = ENUMERATION OF
	(AREA
	,CURVE);
END_TYPE;`,
	"IFCPROJECTEDORTRUELENGTHENUM": `TYPE IfcProjectedOrTrueLengthEnum = ENUMERATION OF
	(PROJECTED_LENGTH
	,TRUE_LENGTH);
END_TYPE;`,
	"IFCPROJECTIONELEMENTTYPEENUM": `TYPE IfcProjectionElementTypeEnum = ENUMERATION OF
	(BLISTER
	,DEVIATOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPROJECTORDERTYPEENUM": `TYPE IfcProjectOrderTypeEnum = ENUMERATION OF
	(CHANGEORDER
	,MAINTENANCEWORKORDER
	,MOVEORDER
	,PURCHASEORDER
	,WORKORDER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPROPERTYSETTEMPLATETYPEENUM": `TYPE IfcPropertySetTemplateTypeEnum = ENUMERATION OF
	(PSET_MATERIALDRIVEN
	,PSET_OCCURRENCEDRIVEN
	,PSET_PERFORMANCEDRIVEN
	,PSET_PROFILEDRIVEN
	,PSET_TYPEDRIVENONLY
	,PSET_TYPEDRIVENOVERRIDE
	,QTO_OCCURRENCEDRIVEN
	,QTO_TYPEDRIVENONLY
	,QTO_TYPEDRIVENOVERRIDE
	,NOTDEFINED);
END_TYPE;`,
	"IFCPROTECTIVEDEVICETRIPPINGUNITTYPEENUM": `TYPE IfcProtectiveDeviceTrippingUnitTypeEnum = ENUMERATION OF
	(ELECTROMAGNETIC
	,ELECTRONIC
	,RESIDUALCURRENT
	,THERMAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPROTECTIVEDEVICETYPEENUM": `TYPE IfcProtectiveDeviceTypeEnum = ENUMERATION OF
	(ANTI_ARCING_DEVICE
	,CIRCUITBREAKER
	,EARTHINGSWITCH
	,EARTHLEAKAGECIRCUITBREAKER
	,FUSEDISCONNECTOR
	,RESIDUALCURRENTCIRCUITBREAKER
	,RESIDUALCURRENTSWITCH
	,SPARKGAP
	,VARISTOR
	,VOLTAGELIMITER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCPUMPTYPEENUM": `TYPE IfcPumpTypeEnum = ENUMERATION OF
	(CIRCULATOR
	,ENDSUCTION
	,SPLITCASE
	,SUBMERSIBLEPUMP
	,SUMPPUMP
	,VERTICALINLINE
	,VERTICALTURBINE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAILINGTYPEENUM": `TYPE IfcRailingTypeEnum = ENUMERATION OF
	(BALUSTRADE
	,FENCE
	,GUARDRAIL
	,HANDRAIL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAILTYPEENUM": `TYPE IfcRailTypeEnum = ENUMERATION OF
	(BLADE
	,CHECKRAIL
	,GUARDRAIL
	,RACKRAIL
	,RAIL
	,STOCKRAIL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAILWAYPARTTYPEENUM": `TYPE IfcRailwayPartTypeEnum = ENUMERATION OF
	(ABOVETRACK
	,DILATIONTRACK
	,LINESIDE
	,LINESIDEPART
	,PLAINTRACK
	,SUBSTRUCTURE
	,TRACK
	,TRACKPART
	,TURNOUTTRACK
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAILWAYTYPEENUM": `TYPE IfcRailwayTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAMPFLIGHTTYPEENUM": `TYPE IfcRampFlightTypeEnum = ENUMERATION OF
	(SPIRAL
	,STRAIGHT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRAMPTYPEENUM": `TYPE IfcRampTypeEnum = ENUMERATION OF
	(HALF_TURN_RAMP
	,QUARTER_TURN_RAMP
	,SPIRAL_RAMP
	,STRAIGHT_RUN_RAMP
	,TWO_QUARTER_TURN_RAMP
	,TWO_STRAIGHT_RUN_RAMP
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCRECURRENCETYPEENUM": `TYPE IfcRecurrenceTypeEnum = ENUMERATION OF
	(BY_DAY_COUNT
	,BY_WEEKDAY_COUNT
	,DAILY
	,MONTHLY_BY_DAY_OF_MONTH
	,MONTHLY_BY_POSITION
	,WEEKLY
	,YEARLY_BY_DAY_OF_MONTH
	,YEARLY_BY_POSITION);
END_TYPE;`,
	"IFCREFERENTTYPEENUM": `TYPE IfcReferentTypeEnum = ENUMERATION OF
	(BOUNDARY
	,INTERSECTION
	,KILOPOINT
	,LANDMARK
	,MILEPOINT
	,POSITION
	,REFERENCEMARKER
	,STATION
	,SUPERELEVATIONEVENT
	,WIDTHEVENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCREFLECTANCEMETHODENUM": `TYPE IfcReflectanceMethodEnum = ENUMERATION OF
	(BLINN
	,FLAT
	,GLASS
	,MATT
	,METAL
	,MIRROR
	,PHONG
	,PHYSICAL
	,PLASTIC
	,STRAUSS
	,NOTDEFINED);
END_TYPE;`,
	"IFCREINFORCEDSOILTYPEENUM": `TYPE IfcReinforcedSoilTypeEnum = ENUMERATION OF
	(DYNAMICALLYCOMPACTED
	,GROUTED
	,REPLACED
	,ROLLERCOMPACTED
	,SURCHARGEPRELOADED
	,VERTICALLYDRAINED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCREINFORCINGBARROLEENUM": `TYPE IfcReinforcingBarRoleEnum = ENUMERATION OF
	(ANCHORING
	,EDGE
	,LIGATURE
	,MAIN
	,PUNCHING
	,RING
	,SHEAR
	,STUD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCREINFORCINGBARSURFACEENUM": `TYPE IfcReinforcingBarSurfaceEnum = ENUMERATION OF
	(PLAIN
	,TEXTURED);
END_TYPE;`,
	"IFCREINFORCINGBARTYPEENUM": `TYPE IfcReinforcingBarTypeEnum = ENUMERATION OF
	(ANCHORING
	,EDGE
	,LIGATURE
	,MAIN
	,PUNCHING
	,RING
	,SHEAR
	,SPACEBAR
	,STUD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCREINFORCINGMESHTYPEENUM": `TYPE IfcReinforcingMeshTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCROADPARTTYPEENUM": `TYPE IfcRoadPartTypeEnum = ENUMERATION OF
	(BICYCLECROSSING
	,BUS_STOP
	,CARRIAGEWAY
	,CENTRALISLAND
	,CENTRALRESERVE
	,HARDSHOULDER
	,INTERSECTION
	,LAYBY
	,PARKINGBAY
	,PASSINGBAY
	,PEDESTRIAN_CROSSING
	,RAILWAYCROSSING
	,REFUGEISLAND
	,ROADSEGMENT
	,ROADSIDE
	,ROADSIDEPART
	,ROADWAYPLATEAU
	,ROUNDABOUT
	,SHOULDER
	,SIDEWALK
	,SOFTSHOULDER
	,TOLLPLAZA
	,TRAFFICISLAND
	,TRAFFICLANE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCROADTYPEENUM": `TYPE IfcRoadTypeEnum = ENUMERATION OF
	(USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCROLEENUM": `TYPE IfcRoleEnum = ENUMERATION OF
	(ARCHITECT
	,BUILDINGOPERATOR
	,BUILDINGOWNER
	,CIVILENGINEER
	,CLIENT
	,COMMISSIONINGENGINEER
	,CONSTRUCTIONMANAGER
	,CONSULTANT
	,CONTRACTOR
	,COSTENGINEER
	,ELECTRICALENGINEER
	,ENGINEER
	,FACILITIESMANAGER
	,FIELDCONSTRUCTIONMANAGER
	,MANUFACTURER
	,MECHANICALENGINEER
	,OWNER
	,PROJECTMANAGER
	,RESELLER
	,STRUCTURALENGINEER
	,SUBCONTRACTOR
	,SUPPLIER
	,USERDEFINED);
END_TYPE;`,
	"IFCROOFTYPEENUM": `TYPE IfcRoofTypeEnum = ENUMERATION OF
	(BARREL_ROOF
	,BUTTERFLY_ROOF
	,DOME_ROOF
	,FLAT_ROOF
	,FREEFORM
	,GABLE_ROOF
	,GAMBREL_ROOF
	,HIPPED_GABLE_ROOF
	,HIP_ROOF
	,MANSARD_ROOF
	,PAVILION_ROOF
	,RAINBOW_ROOF
	,SHED_ROOF
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSANITARYTERMINALTYPEENUM": `TYPE IfcSanitaryTerminalTypeEnum = ENUMERATION OF
	(BATH
	,BIDET
	,CISTERN
	,SANITARYFOUNTAIN
	,SHOWER
	,SINK
	,TOILETPAN
	,URINAL
	,WASHHANDBASIN
	,WCSEAT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSECTIONTYPEENUM": `TYPE IfcSectionTypeEnum = ENUMERATION OF
	(TAPERED
	,UNIFORM);
END_TYPE;`,
	"IFCSENSORTYPEENUM": `TYPE IfcSensorTypeEnum = ENUMERATION OF
	(CO2SENSOR
	,CONDUCTANCESENSOR
	,CONTACTSENSOR
	,COSENSOR
	,EARTHQUAKESENSOR
	,FIRESENSOR
	,FLOWSENSOR
	,FOREIGNOBJECTDETECTIONSENSOR
	,FROSTSENSOR
	,GASSENSOR
	,HEATSENSOR
	,HUMIDITYSENSOR
	,IDENTIFIERSENSOR
	,IONCONCENTRATIONSENSOR
	,LEVELSENSOR
	,LIGHTSENSOR
	,MOISTURESENSOR
	,MOVEMENTSENSOR
	,OBSTACLESENSOR
	,PHSENSOR
	,PRESSURESENSOR
	,RADIATIONSENSOR
	,RADIOACTIVITYSENSOR
	,RAINSENSOR
	,SMOKESENSOR
	,SNOWDEPTHSENSOR
	,SOUNDSENSOR
	,TEMPERATURESENSOR
	,TRAINSENSOR
	,TURNOUTCLOSURESENSOR
	,WHEELSENSOR
	,WINDSENSOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSEQUENCEENUM": `TYPE IfcSequenceEnum = ENUMERATION OF
	(FINISH_FINISH
	,FINISH_START
	,START_FINISH
	,START_START
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSHADINGDEVICETYPEENUM": `TYPE IfcShadingDeviceTypeEnum = ENUMERATION OF
	(AWNING
	,JALOUSIE
	,SHUTTER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSIGNALTYPEENUM": `TYPE IfcSignalTypeEnum = ENUMERATION OF
	(AUDIO
	,MIXED
	,VISUAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSIGNTYPEENUM": `TYPE IfcSignTypeEnum = ENUMERATION OF
	(MARKER
	,MIRROR
	,PICTORAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSIMPLEPROPERTYTEMPLATETYPEENUM": `TYPE IfcSimplePropertyTemplateTypeEnum = ENUMERATION OF
	(P_BOUNDEDVALUE
	,P_ENUMERATEDVALUE
	,P_LISTVALUE
	,P_REFERENCEVALUE
	,P_SINGLEVALUE
	,P_TABLEVALUE
	,Q_AREA
	,Q_COUNT
	,Q_LENGTH
	,Q_NUMBER
	,Q_TIME
	,Q_VOLUME
	,Q_WEIGHT);
END_TYPE;`,
	"IFCSIPREFIX": `TYPE IfcSIPrefix = ENUMERATION OF
	(ATTO
	,CENTI
	,DECA
	,DECI
	,EXA
	,FEMTO
	,GIGA
	,HECTO
	,KILO
	,MEGA
	,MICRO
	,MILLI
	,NANO
	,PETA
	,PICO
	,TERA);
END_TYPE;`,
	"IFCSIUNITNAME": `TYPE IfcSIUnitName = ENUMERATION OF
	(AMPERE
	,BECQUEREL
	,CANDELA
	,COULOMB
	,CUBIC_METRE
	,DEGREE_CELSIUS
	,FARAD
	,GRAM
	,GRAY
	,HENRY
	,HERTZ
	,JOULE
	,KELVIN
	,LUMEN
	,LUX
	,METRE
	,MOLE
	,NEWTON
	,OHM
	,PASCAL
	,RADIAN
	,SECOND
	,SIEMENS
	,SIEVERT
	,SQUARE_METRE
	,STERADIAN
	,TESLA
	,VOLT
	,WATT
	,WEBER);
END_TYPE;`,
	"IFCSLABTYPEENUM": `TYPE IfcSlabTypeEnum = ENUMERATION OF
	(APPROACH_SLAB
	,BASESLAB
	,FLOOR
	,LANDING
	,PAVING
	,ROOF
	,SIDEWALK
	,TRACKSLAB
	,WEARING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSOLARDEVICETYPEENUM": `TYPE IfcSolarDeviceTypeEnum = ENUMERATION OF
	(SOLARCOLLECTOR
	,SOLARPANEL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSPACEHEATERTYPEENUM": `TYPE IfcSpaceHeaterTypeEnum = ENUMERATION OF
	(CONVECTOR
	,RADIATOR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSPACETYPEENUM": `TYPE IfcSpaceTypeEnum = ENUMERATION OF
	(BERTH
	,EXTERNAL
	,GFA
	,INTERNAL
	,PARKING
	,SPACE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSPATIALZONETYPEENUM": `TYPE IfcSpatialZoneTypeEnum = ENUMERATION OF
	(CONSTRUCTION
	,FIRESAFETY
	,INTERFERENCE
	,LIGHTING
	,OCCUPANCY
	,RESERVATION
	,SECURITY
	,THERMAL
	,TRANSPORT
	,VENTILATION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTACKTERMINALTYPEENUM": `TYPE IfcStackTerminalTypeEnum = ENUMERATION OF
	(BIRDCAGE
	,COWL
	,RAINWATERHOPPER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTAIRFLIGHTTYPEENUM": `TYPE IfcStairFlightTypeEnum = ENUMERATION OF
	(CURVED
	,FREEFORM
	,SPIRAL
	,STRAIGHT
	,WINDER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTAIRTYPEENUM": `TYPE IfcStairTypeEnum = ENUMERATION OF
	(CURVED_RUN_STAIR
	,DOUBLE_RETURN_STAIR
	,HALF_TURN_STAIR
	,HALF_WINDING_STAIR
	,LADDER
	,QUARTER_TURN_STAIR
	,QUARTER_WINDING_STAIR
	,SPIRAL_STAIR
	,STRAIGHT_RUN_STAIR
	,THREE_QUARTER_TURN_STAIR
	,THREE_QUARTER_WINDING_STAIR
	,TWO_CURVED_RUN_STAIR
	,TWO_QUARTER_TURN_STAIR
	,TWO_QUARTER_WINDING_STAIR
	,TWO_STRAIGHT_RUN_STAIR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTATEENUM": `TYPE IfcStateEnum = ENUMERATION OF
	(LOCKED
	,READONLY
	,READONLYLOCKED
	,READWRITE
	,READWRITELOCKED);
END_TYPE;`,
	"IFCSTRUCTURALCURVEACTIVITYTYPEENUM": `TYPE IfcStructuralCurveActivityTypeEnum = ENUMERATION OF
	(CONST
	,DISCRETE
	,EQUIDISTANT
	,LINEAR
	,PARABOLA
	,POLYGONAL
	,SINUS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTRUCTURALCURVEMEMBERTYPEENUM": `TYPE IfcStructuralCurveMemberTypeEnum = ENUMERATION OF
	(CABLE
	,COMPRESSION_MEMBER
	,PIN_JOINED_MEMBER
	,RIGID_JOINED_MEMBER
	,TENSION_MEMBER
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTRUCTURALSURFACEACTIVITYTYPEENUM": `TYPE IfcStructuralSurfaceActivityTypeEnum = ENUMERATION OF
	(BILINEAR
	,CONST
	,DISCRETE
	,ISOCONTOUR
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSTRUCTURALSURFACEMEMBERTYPEENUM": `TYPE IfcStructuralSurfaceMemberTypeEnum = ENUMERATION OF
	(BENDING_ELEMENT
	,MEMBRANE_ELEMENT
	,SHELL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSUBCONTRACTRESOURCETYPEENUM": `TYPE IfcSubContractResourceTypeEnum = ENUMERATION OF
	(PURCHASE
	,WORK
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSURFACEFEATURETYPEENUM": `TYPE IfcSurfaceFeatureTypeEnum = ENUMERATION OF
	(DEFECT
	,HATCHMARKING
	,LINEMARKING
	,MARK
	,NONSKIDSURFACING
	,PAVEMENTSURFACEMARKING
	,RUMBLESTRIP
	,SYMBOLMARKING
	,TAG
	,TRANSVERSERUMBLESTRIP
	,TREATMENT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSURFACESIDE": `TYPE IfcSurfaceSide = ENUMERATION OF
	(BOTH
	,NEGATIVE
	,POSITIVE);
END_TYPE;`,
	"IFCSWITCHINGDEVICETYPEENUM": `TYPE IfcSwitchingDeviceTypeEnum = ENUMERATION OF
	(CONTACTOR
	,DIMMERSWITCH
	,EMERGENCYSTOP
	,KEYPAD
	,MOMENTARYSWITCH
	,RELAY
	,SELECTORSWITCH
	,STARTER
	,START_AND_STOP_EQUIPMENT
	,SWITCHDISCONNECTOR
	,TOGGLESWITCH
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCSYSTEMFURNITUREELEMENTTYPEENUM": `TYPE IfcSystemFurnitureElementTypeEnum = ENUMERATION OF
	(PANEL
	,SUBRACK
	,WORKSURFACE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTANKTYPEENUM": `TYPE IfcTankTypeEnum = ENUMERATION OF
	(BASIN
	,BREAKPRESSURE
	,EXPANSION
	,FEEDANDEXPANSION
	,OILRETENTIONTRAY
	,PRESSUREVESSEL
	,STORAGE
	,VESSEL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTASKDURATIONENUM": `TYPE IfcTaskDurationEnum = ENUMERATION OF
	(ELAPSEDTIME
	,WORKTIME
	,NOTDEFINED);
END_TYPE;`,
	"IFCTASKTYPEENUM": `TYPE IfcTaskTypeEnum = ENUMERATION OF
	(ADJUSTMENT
	,ATTENDANCE
	,CALIBRATION
	,CONSTRUCTION
	,DEMOLITION
	,DISMANTLE
	,DISPOSAL
	,EMERGENCY
	,INSPECTION
	,INSTALLATION
	,LOGISTIC
	,MAINTENANCE
	,MOVE
	,OPERATION
	,REMOVAL
	,RENOVATION
	,SAFETY
	,SHUTDOWN
	,STARTUP
	,TESTING
	,TROUBLESHOOTING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTENDONANCHORTYPEENUM": `TYPE IfcTendonAnchorTypeEnum = ENUMERATION OF
	(COUPLER
	,FIXED_END
	,TENSIONING_END
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTENDONCONDUITTYPEENUM": `TYPE IfcTendonConduitTypeEnum = ENUMERATION OF
	(COUPLER
	,DIABOLO
	,DUCT
	,GROUTING_DUCT
	,TRUMPET
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTENDONTYPEENUM": `TYPE IfcTendonTypeEnum = ENUMERATION OF
	(BAR
	,COATED
	,STRAND
	,WIRE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTEXTPATH": `TYPE IfcTextPath = ENUMERATION OF
	(DOWN
	,LEFT
	,RIGHT
	,UP);
END_TYPE;`,
	"IFCTIMESERIESDATATYPEENUM": `TYPE IfcTimeSeriesDataTypeEnum = ENUMERATION OF
	(CONTINUOUS
	,DISCRETE
	,DISCRETEBINARY
	,PIECEWISEBINARY
	,PIECEWISECONSTANT
	,PIECEWISECONTINUOUS
	,NOTDEFINED);
END_TYPE;`,
	"IFCTRACKELEMENTTYPEENUM": `TYPE IfcTrackElementTypeEnum = ENUMERATION OF
	(BLOCKINGDEVICE
	,DERAILER
	,FROG
	,HALF_SET_OF_BLADES
	,SLEEPER
	,SPEEDREGULATOR
	,TRACKENDOFALIGNMENT
	,VEHICLESTOP
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTRANSFORMERTYPEENUM": `TYPE IfcTransformerTypeEnum = ENUMERATION OF
	(CHOPPER
	,COMBINED
	,CURRENT
	,FREQUENCY
	,INVERTER
	,RECTIFIER
	,VOLTAGE
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTRANSITIONCODE": `TYPE IfcTransitionCode = ENUMERATION OF
	(CONTINUOUS
	,CONTSAMEGRADIENT
	,CONTSAMEGRADIENTSAMECURVATURE
	,DISCONTINUOUS);
END_TYPE;`,
	"IFCTRANSPORTELEMENTTYPEENUM": `TYPE IfcTransportElementTypeEnum = ENUMERATION OF
	(CRANEWAY
	,ELEVATOR
	,ESCALATOR
	,HAULINGGEAR
	,LIFTINGGEAR
	,MOVINGWALKWAY
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCTRIMMINGPREFERENCE": `TYPE IfcTrimmingPreference = ENUMERATION OF
	(CARTESIAN
	,PARAMETER
	,UNSPECIFIED);
END_TYPE;`,
	"IFCTUBEBUNDLETYPEENUM": `TYPE IfcTubeBundleTypeEnum = ENUMERATION OF
	(FINNED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCUNITARYCONTROLELEMENTTYPEENUM": `TYPE IfcUnitaryControlElementTypeEnum = ENUMERATION OF
	(ALARMPANEL
	,BASESTATIONCONTROLLER
	,COMBINED
	,CONTROLPANEL
	,GASDETECTIONPANEL
	,HUMIDISTAT
	,INDICATORPANEL
	,MIMICPANEL
	,THERMOSTAT
	,WEATHERSTATION
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCUNITARYEQUIPMENTTYPEENUM": `TYPE IfcUnitaryEquipmentTypeEnum = ENUMERATION OF
	(AIRCONDITIONINGUNIT
	,AIRHANDLER
	,DEHUMIDIFIER
	,ROOFTOPUNIT
	,SPLITSYSTEM
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCUNITENUM": `TYPE IfcUnitEnum = ENUMERATION OF
	(ABSORBEDDOSEUNIT
	,AMOUNTOFSUBSTANCEUNIT
	,AREAUNIT
	,DOSEEQUIVALENTUNIT
	,ELECTRICCAPACITANCEUNIT
	,ELECTRICCHARGEUNIT
	,ELECTRICCONDUCTANCEUNIT
	,ELECTRICCURRENTUNIT
	,ELECTRICRESISTANCEUNIT
	,ELECTRICVOLTAGEUNIT
	,ENERGYUNIT
	,FORCEUNIT
	,FREQUENCYUNIT
	,ILLUMINANCEUNIT
	,INDUCTANCEUNIT
	,LENGTHUNIT
	,LUMINOUSFLUXUNIT
	,LUMINOUSINTENSITYUNIT
	,MAGNETICFLUXDENSITYUNIT
	,MAGNETICFLUXUNIT
	,MASSUNIT
	,PLANEANGLEUNIT
	,POWERUNIT
	,PRESSUREUNIT
	,RADIOACTIVITYUNIT
	,SOLIDANGLEUNIT
	,THERMODYNAMICTEMPERATUREUNIT
	,TIMEUNIT
	,VOLUMEUNIT
	,USERDEFINED);
END_TYPE;`,
	"IFCVALVETYPEENUM": `TYPE IfcValveTypeEnum = ENUMERATION OF
	(AIRRELEASE
	,ANTIVACUUM
	,CHANGEOVER
	,CHECK
	,COMMISSIONING
	,DIVERTING
	,DOUBLECHECK
	,DOUBLEREGULATING
	,DRAWOFFCOCK
	,FAUCET
	,FLUSHING
	,GASCOCK
	,GASTAP
	,ISOLATING
	,MIXING
	,PRESSUREREDUCING
	,PRESSURERELIEF
	,REGULATING
	,SAFETYCUTOFF
	,STEAMTRAP
	,STOPCOCK
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCVEHICLETYPEENUM": `TYPE IfcVehicleTypeEnum = ENUMERATION OF
	(CARGO
	,ROLLINGSTOCK
	,VEHICLE
	,VEHICLEAIR
	,VEHICLEMARINE
	,VEHICLETRACKED
	,VEHICLEWHEELED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCVIBRATIONDAMPERTYPEENUM": `TYPE IfcVibrationDamperTypeEnum = ENUMERATION OF
	(AXIAL_YIELD
	,BENDING_YIELD
	,FRICTION
	,RUBBER
	,SHEAR_YIELD
	,VISCOUS
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCVIBRATIONISOLATORTYPEENUM": `TYPE IfcVibrationIsolatorTypeEnum = ENUMERATION OF
	(BASE
	,COMPRESSION
	,SPRING
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCVIRTUALELEMENTTYPEENUM": `TYPE IfcVirtualElementTypeEnum = ENUMERATION OF
	(BOUNDARY
	,CLEARANCE
	,PROVISIONFORVOID
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCVOIDINGFEATURETYPEENUM": `TYPE IfcVoidingFeatureTypeEnum = ENUMERATION OF
	(CHAMFER
	,CUTOUT
	,EDGE
	,HOLE
	,MITER
	,NOTCH
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWALLTYPEENUM": `TYPE IfcWallTypeEnum = ENUMERATION OF
	(ELEMENTEDWALL
	,MOVABLE
	,PARAPET
	,PARTITIONING
	,PLUMBINGWALL
	,POLYGONAL
	,RETAININGWALL
	,SHEAR
	,SOLIDWALL
	,STANDARD
	,WAVEWALL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWASTETERMINALTYPEENUM": `TYPE IfcWasteTerminalTypeEnum = ENUMERATION OF
	(FLOORTRAP
	,FLOORWASTE
	,GULLYSUMP
	,GULLYTRAP
	,ROOFDRAIN
	,WASTEDISPOSALUNIT
	,WASTETRAP
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWINDOWPANELOPERATIONENUM": `TYPE IfcWindowPanelOperationEnum = ENUMERATION OF
	(BOTTOMHUNG
	,FIXEDCASEMENT
	,OTHEROPERATION
	,PIVOTHORIZONTAL
	,PIVOTVERTICAL
	,REMOVABLECASEMENT
	,SIDEHUNGLEFTHAND
	,SIDEHUNGRIGHTHAND
	,SLIDINGHORIZONTAL
	,SLIDINGVERTICAL
	,TILTANDTURNLEFTHAND
	,TILTANDTURNRIGHTHAND
	,TOPHUNG
	,NOTDEFINED);
END_TYPE;`,
	"IFCWINDOWPANELPOSITIONENUM": `TYPE IfcWindowPanelPositionEnum = ENUMERATION OF
	(BOTTOM
	,LEFT
	,MIDDLE
	,RIGHT
	,TOP
	,NOTDEFINED);
END_TYPE;`,
	"IFCWINDOWTYPEENUM": `TYPE IfcWindowTypeEnum = ENUMERATION OF
	(LIGHTDOME
	,SKYLIGHT
	,WINDOW
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWINDOWTYPEPARTITIONINGENUM": `TYPE IfcWindowTypePartitioningEnum = ENUMERATION OF
	(DOUBLE_PANEL_HORIZONTAL
	,DOUBLE_PANEL_VERTICAL
	,SINGLE_PANEL
	,TRIPLE_PANEL_BOTTOM
	,TRIPLE_PANEL_HORIZONTAL
	,TRIPLE_PANEL_LEFT
	,TRIPLE_PANEL_RIGHT
	,TRIPLE_PANEL_TOP
	,TRIPLE_PANEL_VERTICAL
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWORKCALENDARTYPEENUM": `TYPE IfcWorkCalendarTypeEnum = ENUMERATION OF
	(FIRSTSHIFT
	,SECONDSHIFT
	,THIRDSHIFT
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWORKPLANTYPEENUM": `TYPE IfcWorkPlanTypeEnum = ENUMERATION OF
	(ACTUAL
	,BASELINE
	,PLANNED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
	"IFCWORKSCHEDULETYPEENUM": `TYPE IfcWorkScheduleTypeEnum = ENUMERATION OF
	(ACTUAL
	,BASELINE
	,PLANNED
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`,
}
